package capture

/*
#include <linux/videodev2.h>
#include <sys/mman.h>
#include <sys/types.h>

int somemacro (int fd) {
	fd_set fds;
	FD_ZERO(&fds);
	FD_SET(fd, &fds);
	struct timeval tv = {0};
	tv.tv_sec = 20;
	return select(fd+1, &fds, 0, 0, &tv);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"image"
	"os"
	"syscall"
	"unsafe"
)

var (
	ErrMJPEGNotSupported = errors.New("this webcam dosen't support Motion-JPEG")
	ErrYUYVNotSupported  = errors.New("this webcam dosen't support YUV 4:2:2")
	ErrInvalidFormat     = errors.New("invalid image format")
	ErrUnknown           = errors.New("unknown error")
)

func Capture(device string) (image.Image, error) {
	vd, err := os.OpenFile(device, os.O_RDWR, 0660)
	if err != nil {
		return nil, err
	}
	defer vd.Close()

	width, height, formats, err := getInfo(vd, true)
	if err != nil {
		return nil, err
	}

	// select YUV 4:2:2 (YUYV)
	pixfmt := C.V4L2_PIX_FMT_YUYV
	ok := false
	for _, v := range formats {
		if v == C.V4L2_PIX_FMT_YUYV {
			ok = true
			break
		}
	}

	if !ok {
		return nil, ErrYUYVNotSupported
	}

	width, height, err = setFormat(vd, width, height, pixfmt)
	if err != nil {
		return nil, err
	}

	b, err := getFrame(vd)
	if err != nil {
		return nil, err
	}

	return toImage(pixfmt, b, width, height), nil
}

func ioctl(fd *os.File, op uintptr, arg uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), op, arg)
	if err != 0 {
		return err
	}
	return nil
}

func getInfo(vd *os.File, show bool) (int, int, []C.__u32, error) {
	// get capability info
	var caps C.struct_v4l2_capability

	err := ioctl(vd, C.VIDIOC_QUERYCAP, uintptr(unsafe.Pointer(&caps)))
	if err != nil {
		return 0, 0, nil, err
	}
	if show {
		fmt.Printf("Driver Capability:\n" +
			"    Driver: \"%s\"\n" +
			"    Card: \"%s\"\n" +
			"    Bus: \"%s\"\n" +
			"    Capabilities: %08x\n")
	}

	// get crop info
	var cropcap C.struct_v4l2_cropcap
	cropcap._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE

	err = ioctl(vd, C.VIDIOC_CROPCAP, uintptr(unsafe.Pointer(&cropcap)))
	if err != nil {
		return 0, 0, nil, err
	}

	if show {
		fmt.Printf("Cropping Capability:\n"+
			"    Bounds: %dx%d+%d+%d\n"+
			"    Default: %dx%d+%d+%d\n"+
			"    Aspect: %d/%d\n",
			cropcap.bounds.width, cropcap.bounds.height, cropcap.bounds.left, cropcap.bounds.top,
			cropcap.defrect.width, cropcap.defrect.height, cropcap.defrect.left, cropcap.defrect.top,
			cropcap.pixelaspect.numerator, cropcap.pixelaspect.denominator)
	}

	// get format info
	var fmtdesc C.struct_v4l2_fmtdesc
	fmtdesc._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE

	pixfmt := make([]C.__u32, 0)
	if show {
		fmt.Println("Format Description:")
	}
	for {
		err = ioctl(vd, C.VIDIOC_ENUM_FMT, uintptr(unsafe.Pointer(&fmtdesc)))
		if err != nil {
			break
		}

		pixfmt = append(pixfmt, fmtdesc.pixelformat)

		if show {
			var dst [32]C.char
			for i, v := range fmtdesc.description {
				dst[i] = C.char(v)
			}
			fmt.Printf("    %s\n", C.GoString((*C.char)(&dst[0])))
		}
		fmtdesc.index++
	}

	return int(cropcap.defrect.width), int(cropcap.defrect.height), pixfmt, nil
}

func setFormat(vd *os.File, w, h, pixfmt int) (int, int, error) {
	var format C.struct_v4l2_format
	format._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE
	// get struct_v4l2_pix_format in union fmt in struct_v4l2_format
	pix := (*C.struct_v4l2_pix_format)(unsafe.Pointer(&format.fmt[0]))
	pix.width = C.__u32(w)
	pix.height = C.__u32(h)
	pix.pixelformat = C.__u32(pixfmt)
	pix.field = C.V4L2_FIELD_NONE

	err := ioctl(vd, C.VIDIOC_S_FMT, uintptr(unsafe.Pointer(&format)))
	if err != nil {
		return 0, 0, err
	} else if pix.pixelformat != C.__u32(pixfmt) {
		return 0, 0, ErrInvalidFormat
	}

	return int(pix.width), int(pix.height), nil
}

func getFrame(vd *os.File) ([]byte, error) {
	// request buffer
	var reqbuf C.struct_v4l2_requestbuffers
	reqbuf._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE
	reqbuf.count = 1
	reqbuf.memory = C.V4L2_MEMORY_MMAP

	err := ioctl(vd, C.VIDIOC_REQBUFS, uintptr(unsafe.Pointer(&reqbuf)))
	if err != nil {
		return nil, err
	}

	// query buffer
	var buf C.struct_v4l2_buffer
	buf._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE
	buf.memory = C.V4L2_MEMORY_MMAP
	buf.index = 0

	err = ioctl(vd, C.VIDIOC_QUERYBUF, uintptr(unsafe.Pointer(&buf)))
	if err != nil {
		return nil, err
	}

	buffer := C.mmap(nil, C.size_t(buf.length), C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, C.int(vd.Fd()),
		*(*C.__off_t)(unsafe.Pointer(&buf.m[0])))
	fmt.Printf("Length: %d\nAddress: %p\n", buf.length, buffer)
	fmt.Printf("Image Length: %d\n", buf.bytesused)

	var qbuf C.struct_v4l2_buffer
	qbuf._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE
	qbuf.memory = C.V4L2_MEMORY_MMAP
	qbuf.index = 0

	err = ioctl(vd, C.VIDIOC_QBUF, uintptr(unsafe.Pointer(&qbuf)))
	if err != nil {
		return nil, err
	}

	err = ioctl(vd, C.VIDIOC_STREAMON, uintptr(unsafe.Pointer(&qbuf._type)))
	if err != nil {
		return nil, err
	}

	e := C.somemacro(C.int(vd.Fd()))
	if e == -1 {
		return nil, ErrUnknown
	}

	err = ioctl(vd, C.VIDIOC_DQBUF, uintptr(unsafe.Pointer(&qbuf)))
	if err != nil {
		return nil, err
	}

	return C.GoBytes(buffer, C.int(qbuf.bytesused)), nil
}

// convert YUYV(YUV422) data to image.Image
func toImage(pixfmt int, b []byte, w, h int) image.Image {
	var ratio image.YCbCrSubsampleRatio
	switch pixfmt {
	case C.V4L2_PIX_FMT_YUYV:
		ratio = image.YCbCrSubsampleRatio422
	}

	// ** yCbCr422 (YUYV) **
	// 4byte = 2pixel
	//
	// 4byte: Y0, Cb, Y1, Cr
	// pixel1: Y0, Cb, Cr
	// pixel2: Y1, Cb, Cr
	pixs := len(b) / 4 * 2

	img := image.NewYCbCr(image.Rect(0, 0, w, h), ratio)
	img.YStride = w
	img.CStride = w / 2
	img.Y = make([]byte, pixs)
	img.Cb = make([]byte, pixs/2)
	img.Cr = make([]byte, pixs/2)

	for i := 0; i < pixs/2; i++ {
		img.Y[i*2] = b[i*4]
		img.Cb[i] = b[i*4+1]
		img.Y[i*2+1] = b[i*4+2]
		img.Cr[i] = b[i*4+3]
	}

	return img
}

package main

/*
#include <linux/videodev2.h>
#include <sys/mman.h>
#include <sys/types.h>

int capture (int fd) {
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
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
)

var (
	ErrMJPEGNotSupported = errors.New("This webcam dosen't support Motion-JPEG")
	ErrFormatInvalid     = errors.New("Invalid image format")
)

func main() {
	vd, err := os.OpenFile("/dev/video0", os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer vd.Close()

	if err = printInfo(vd); err != nil {
		fmt.Println(err)
		return
	}
	if err = setMode(vd); err != nil {
		fmt.Println(err)
		return
	}
	if err = save(vd); err != nil {
		fmt.Println(err)
		return
	}
}

func ioctl(fd *os.File, op uintptr, arg uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), op, arg)
	if err != 0 {
		return err
	}
	return nil
}

func printInfo(vd *os.File) error {
	// get capability info
	var caps C.struct_v4l2_capability

	err := ioctl(vd, C.VIDIOC_QUERYCAP, uintptr(unsafe.Pointer(&caps)))
	if err != nil {
		return err
	}

	fmt.Printf("Driver Capability:\n"+
		"    Driver: \"%s\"\n"+
		"    Card: \"%s\"\n"+
		"    Bus: \"%s\"\n"+
		"    Capabilities: %08x\n",
		caps.driver,
		caps.card,
		caps.bus_info,
		caps.capabilities)

	// get crop info
	var cropcap C.struct_v4l2_cropcap
	cropcap._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE

	err = ioctl(vd, C.VIDIOC_CROPCAP, uintptr(unsafe.Pointer(&cropcap)))
	if err != nil {
		return err
	}

	fmt.Printf("Cropping Capability:\n"+
		"    Bounds: %dx%d+%d+%d\n"+
		"    Default: %dx%d+%d+%d\n"+
		"    Aspect: %d/%d\n",
		cropcap.bounds.width, cropcap.bounds.height, cropcap.bounds.left, cropcap.bounds.top,
		cropcap.defrect.width, cropcap.defrect.height, cropcap.defrect.left, cropcap.defrect.top,
		cropcap.pixelaspect.numerator, cropcap.pixelaspect.denominator)

	// get format info
	var fmtdesc C.struct_v4l2_fmtdesc
	fmtdesc._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE

	fmt.Println("Format Description:")
	for {
		err = ioctl(vd, C.VIDIOC_ENUM_FMT, uintptr(unsafe.Pointer(&fmtdesc)))
		if err != nil {
			break
		}
		var dst [32]C.char
		for i, v := range fmtdesc.description {
			dst[i] = C.char(v)
		}
		fmt.Printf("    %s\n", C.GoString((*C.char)(&dst[0])))
		fmtdesc.index++
	}

	return nil
}

func setMode(vd *os.File) error {
	// check the webcam supports Motion-JPEG
	var fmtdesc C.struct_v4l2_fmtdesc
	fmtdesc._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE

	mjpeg := false
	for {
		err := ioctl(vd, C.VIDIOC_ENUM_FMT, uintptr(unsafe.Pointer(&fmtdesc)))
		if err != nil {
			break
		}
		if fmtdesc.pixelformat == C.V4L2_PIX_FMT_MJPEG {
			mjpeg = true
			break
		}
		fmtdesc.index++
	}

	if !mjpeg {
		return ErrMJPEGNotSupported
	}

	// set format
	var format C.struct_v4l2_format
	format._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE
	// get struct_v4l2_pix_format in union fmt in struct_v4l2_format
	pix := (*C.struct_v4l2_pix_format)(unsafe.Pointer(&format.fmt[0]))
	//pix := C.getpix(&format)
	pix.width = 640
	pix.height = 480
	pix.pixelformat = C.V4L2_PIX_FMT_MJPEG
	pix.field = C.V4L2_FIELD_NONE

	err := ioctl(vd, C.VIDIOC_S_FMT, uintptr(unsafe.Pointer(&format)))
	if err != nil || pix.pixelformat != C.V4L2_PIX_FMT_MJPEG {
		return ErrFormatInvalid
	}

	fmt.Printf("Selected Mode:\n"+
		"    Width: %d\n"+
		"    Height: %d\n"+
		"    Bytes per Line: %d\n"+
		"    Pixel Format: %s\n"+
		"    Field: %d\n",
		pix.width,
		pix.height,
		pix.bytesperline,
		"MJPEG",
		pix.field)

	/*
		// set jpeg compression
		var jpegcomp C.struct_v4l2_jpegcompression
		err = ioctl(vd, C.VIDIOC_G_JPEGCOMP, uintptr(unsafe.Pointer(&jpegcomp)))
		if err != nil {
			return err
		}


			fmt.Printf("JPEG Compression:\n"+
					"    Quality: %d\n"+
					"    APPn: %d\n"+
					"    APP_len: %d\n"+
					"    COM_len: %d\n"+
					"    DHT: %b\n"+
					"    DQT: %b\n"+
					"    DRI: %b\n"+
					"    COM: %b\n"+
					"    APP: %b\n",
					jpegcomp.quality,
					jpegcomp.APPn,
					jpegcomp.APP_len,
					jpegcomp.COM_len,
					jpegcomp.jpeg_markers|C.V4L2_JPEG_MARKER_DHT != 0,
					jpegcomp.jpeg_markers|C.V4L2_JPEG_MARKER_DQT != 0,
					jpegcomp.jpeg_markers|C.V4L2_JPEG_MARKER_DRI != 0,
					jpegcomp.jpeg_markers|C.V4L2_JPEG_MARKER_COM != 0,
					jpegcomp.jpeg_markers|C.V4L2_JPEG_MARKER_APP != 0)
	*/

	return nil
}

func save(vd *os.File) error {
	// request buffer
	var reqbuf C.struct_v4l2_requestbuffers
	reqbuf._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE
	reqbuf.count = 1
	reqbuf.memory = C.V4L2_MEMORY_MMAP

	err := ioctl(vd, C.VIDIOC_REQBUFS, uintptr(unsafe.Pointer(&reqbuf)))
	if err != nil {
		return err
	}

	// query buffer
	var buf C.struct_v4l2_buffer
	buf._type = C.V4L2_BUF_TYPE_VIDEO_CAPTURE
	buf.memory = C.V4L2_MEMORY_MMAP
	buf.index = 0

	err = ioctl(vd, C.VIDIOC_QUERYBUF, uintptr(unsafe.Pointer(&buf)))
	if err != nil {
		return err
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
		return err
	}

	err = ioctl(vd, C.VIDIOC_STREAMON, uintptr(unsafe.Pointer(&qbuf._type)))
	if err != nil {
		return err
	}

	e := C.capture(C.int(vd.Fd()))
	if e == -1 {
		fmt.Printf("Error")
		return nil
	}

	err = ioctl(vd, C.VIDIOC_DQBUF, uintptr(unsafe.Pointer(&qbuf)))
	if err != nil {
		return err
	}

	return ioutil.WriteFile("test.jpg", C.GoBytes(buffer, C.int(qbuf.bytesused)), 0644)
}

package main

import (
	. "github.com/OUCC/prism/logger"

	cv "github.com/lazywei/go-opencv/opencv"

	"bytes"
	"encoding/base64"
	"encoding/json"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	cam0 *cv.Capture
)

type imageData struct {
	Key   string `json:"key"`
	Image string `json:"image"`
}

func main() {
	setupCamera()

	for {
		if img := capture(); img != nil {
			post(img)
		}

		time.Sleep(CAPTURE_INTERVAL)
	}
}

func setupCamera() {
	cam0 = cv.NewCameraCapture(0)
	if cam0 == nil {
		Log.Fatal("can not open camera")
	}
}

func capture() []byte {
	if !cam0.GrabFrame() {
		return nil
	}

	//!!!DO NOT RELEASE or MODIFY the retrieved frame!!!
	img := cam0.RetrieveFrame(1)
	if img == nil {
		Log.Error("failed to capture")
		return nil
	}

	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img.ToImage(), nil)
	if err != nil {
		Log.Error(err.Error())
		return nil
	}
	if DEBUG {
		ioutil.WriteFile("test.jpg", buf.Bytes(), 0644)
	}
	return buf.Bytes()
}

func post(img []byte) {
	b, _ := json.Marshal(imageData{
		Key:   PRISM_KEY,
		Image: base64.StdEncoding.EncodeToString(img),
	})
	resp, err := http.Post(IMG_POST_URL, "application/json", bytes.NewReader(b))
	if err != nil {
		Log.Error(err.Error())
	} else {
		Log.Info(resp.Status)
	}
}

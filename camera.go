package main

import (
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

func init() {
	cam0 = cv.NewCameraCapture(0)
	if cam0 == nil {
		Log.Fatal("can not open camera")
	}
}

func captureAndPost() {
	for {
		if DEBUG {
			time.Sleep(30 * time.Second)
		} else {
			time.Sleep(10 * time.Minute)
		}

		if !cam0.GrabFrame() {
			continue
		}

		//!!!DO NOT RELEASE or MODIFY the retrieved frame!!!
		img := cam0.RetrieveFrame(1)
		if img == nil {
			Log.Error("failed to capture")
			continue
		}

		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, img.ToImage(), nil)
		if err != nil {
			Log.Error(err.Error())
			return
		}
		if DEBUG {
			ioutil.WriteFile("test.jpg", buf.Bytes(), 0644)
		}

		b, _ := json.Marshal(imageData{
			Key:   PRISM_KEY,
			Image: base64.StdEncoding.EncodeToString(buf.Bytes()),
		})
		resp, err := http.Post(IMG_POST_URL, "application/json", bytes.NewReader(b))
		if err != nil {
			Log.Error(err.Error())
		} else {
			Log.Info(resp.Status)
		}
	}
}

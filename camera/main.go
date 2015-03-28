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
	"net/url"
	"time"
)

func main() {
	for {
		if img := capture(); img != nil {
			post(img)
			post2(img)
		}

		time.Sleep(CAPTURE_INTERVAL)
	}
}

func capture() []byte {
	cam := cv.NewCameraCapture(0)
	if cam == nil {
		Log.Error("can not open camera")
		return nil
	}
	defer cam.Release()

	for i := 0; i < DROP_FRAME; i++ {
		cam.QueryFrame()
	}

	//!!!DO NOT RELEASE or MODIFY the retrieved frame!!!
	img := cam.QueryFrame()
	if img == nil {
		Log.Error("failed to capture")
		return nil
	}

	// convert IplImage to jpeg
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
	type post struct {
		Key   string `json:"key"`
		Image string `json:"image"`
	}

	b, _ := json.Marshal(post{
		Key:   PRISM_KEY,
		Image: base64.StdEncoding.EncodeToString(img),
	})
	resp, err := http.Post(IMG_POST_URL, "application/json", bytes.NewReader(b))
	if err != nil {
		Log.Error(err.Error())
	} else if resp.StatusCode != http.StatusCreated {
		Log.Error(resp.Status)
	} else {
		Log.Info(resp.Status)
	}
}

func post2(img []byte) {
	values := url.Values{}
	values.Add("image", base64.StdEncoding.EncodeToString(img))
	values.Add("key", PRISM_KEY)
	resp, err := http.PostForm(IMG_POST_URL2, values)
	if err != nil {
		Log.Error(err.Error())
	} else if resp.StatusCode != http.StatusOK {
		Log.Error(resp.Status)
	} else {
		Log.Info(resp.Status)
	}
}

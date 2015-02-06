package main

import (
	"github.com/OUCC/prism/capture"

	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	//"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	VIDEO_DEVICE = "/dev/video0"
)

func main() {
	for {
		img, err := capture.Capture(VIDEO_DEVICE)
		if err != nil {
			fmt.Println("Error: failed to capture: ", err)
			return
		}

		buf := new(bytes.Buffer)
		err = jpeg.Encode(buf, img, nil)
		if err != nil {
			fmt.Println(err)
		}
		//ioutil.WriteFile("test.jpg", buf.Bytes(), 0644) // for debug

		s := base64.StdEncoding.EncodeToString(buf.Bytes())
		if err := post(s); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(time.Now().Format(time.UnixDate), "StatusOK")
		}

		time.Sleep(5 * time.Minute)
	}
}

func post(img string) error {
	values := url.Values{}
	values.Add("image", img)
	values.Add("key", POST_KEY)
	res, err := http.PostForm(POST_URL, values)
	if err != nil || res.StatusCode != http.StatusOK {
		return err
	}
	return nil
}

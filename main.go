package main

import (
	"github.com/OUCC/prism/capture"

	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	VIDEO_DEVICE = "/dev/video0"
)

func main() {
	img, err := capture.Capture(VIDEO_DEVICE)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("test.jpg", buf.Bytes(), 0644)
	return

	s := base64.StdEncoding.EncodeToString(buf.Bytes())
	if err := post(s); err != nil {
		fmt.Println(err)
	}
}

func post(img string) error {
	var values url.Values
	values.Add("image", img)
	values.Add("key", POST_KEY)
	res, err := http.PostForm(POST_URL, values)
	if err != nil {
		return err
	}
	fmt.Printf("Status: %d\n", res.StatusCode)
	// TOOD
	return nil
}

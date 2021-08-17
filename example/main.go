package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/shanghuiyang/imgr"
	"github.com/shanghuiyang/oauth"
)

const (
	apiKey    = "your_baidu_api_key"
	secretKey = "your_baidu_secret_key"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: invalid args")
		fmt.Println("usage: image-recognizer test.jpg")
		os.Exit(1)
	}
	imgf := os.Args[1]
	img, err := ioutil.ReadFile(imgf)
	if err != nil {
		log.Printf("failed to read image file: %v, error: %v", imgf, err)
		os.Exit(1)
	}

	auth := oauth.NewBaiduOauth(apiKey, secretKey, oauth.NewCacheImp())
	r := imgr.NewBaiduRecognizer(auth)
	text, err := r.Recognize(img)
	if err != nil {
		log.Printf("failed to recognize the image, error: %v", err)
		os.Exit(1)
	}
	fmt.Println(text)
}

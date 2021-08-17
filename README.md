# imgr
imgr recognizes the object in an image using the APIs like baidu and google.

## Usage
```go
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
    imgf := "images/dog.jpg"
	img, err := ioutil.ReadFile(imgf)
	if err != nil {
		fmt.Printf("failed to read image file: %v, error: %v\n", imgf, err)
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
```

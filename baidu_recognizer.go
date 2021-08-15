package imgr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/shanghuiyang/oauth"
)

const (
	baiduURL = "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general"
)

// Recognizer ...
type BaiduRecognizer struct {
	auth oauth.Oauth
}

type response struct {
	ResultNum int32    `json:"result_num"`
	Results   []result `json:"result"`
	ErrorCode int      `json:"error_code"`
	ErrorMsg  string   `json:"error_msg"`
}

type result struct {
	Score   float32 `json:"score"`
	Root    string  `json:"root"`
	Keyword string  `json:"keyword"`
}

// New ...
func NewBaiduRecognizer(auth oauth.Oauth) *BaiduRecognizer {
	return &BaiduRecognizer{
		auth: auth,
	}
}

// Recognize ...
func (r *BaiduRecognizer) Recognize(image []byte) (string, error) {
	token, err := r.auth.Token()
	if err != nil {
		return "", err
	}

	b64img := b64Image(image)
	formData := url.Values{
		"access_token": {token},
		"image":        {b64img},
	}
	resp, err := http.PostForm(baiduURL, formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var res response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}
	if res.ErrorCode > 0 {
		return "", fmt.Errorf("error_code: %v, error_msg: %v", res.ErrorCode, res.ErrorMsg)
	}

	if res.ResultNum > 0 {
		return res.Results[0].Keyword, nil
	}
	return "", fmt.Errorf("failed to recognize")
}

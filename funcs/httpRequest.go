package funcs

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/yzbtdiy/alist_batch/models"
)

// reference: https://golangtutorial.dev/tips/http-post-json-go/
// 封装 post 请求
func HttpPost(url string, token string, postData []byte) models.ResData {
	// log.Println("HTTP JSON POST URL:", url)
	// var jsonData = []byte(postData)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if token != "" {
		request.Header.Set("Authorization", token)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	var resData models.ResData
	err = json.Unmarshal(body, &resData)
	if err != nil {
		panic(err)
	}

	return resData
}

// 封装 get 请求
func HttpGet(url string, token string) models.ResData {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	// req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var resData models.ResData
	err = json.Unmarshal(body, &resData)
	if err != nil {
		panic(err)
	}
	return resData
}

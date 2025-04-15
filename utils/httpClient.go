package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/yzbtdiy/alist_batch/models"
)

// HttpClient 封装 http.Client 和 token
type HttpClient struct {
	client *http.Client
	token  string
}

// 创建新的 HttpClient 实例
func NewHttpClient(token string) *HttpClient {
	client := new(http.Client)

	return &HttpClient{
		client: client,
		token:  token,
	}
}

// Post 封装 POST 请求
func (hc *HttpClient) Post(url string, postData []byte) models.ResData {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if hc.token != "" {
		request.Header.Set("Authorization", hc.token)
	}

	response, err := hc.client.Do(request)
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

// Get 封装 GET 请求
func (hc *HttpClient) Get(url string) models.ResData {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	if hc.token != "" {
		req.Header.Set("Authorization", hc.token)
	}

	res, err := hc.client.Do(req)
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

// Close 关闭底层连接（如支持）
func (hc *HttpClient) Close() {
	if closer, ok := hc.client.Transport.(interface{ Close() error }); ok {
		_ = closer.Close()
	}
}

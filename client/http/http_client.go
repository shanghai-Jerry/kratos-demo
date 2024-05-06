package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	*http.Client
	URL string
}

func NewHTTPClient(URL string) *Client {
	c := &Client{
		Client: &http.Client{},
		URL:    URL,
	}
	return c
}

// do 方法发送HTTP请求并返回响应状态码、响应体和错误信息
func (c *Client) do(req *http.Request) (int, []byte, error) {
	resp, err := c.Do(req)
	if err != nil {
		return http.StatusServiceUnavailable, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, body, nil
}

func (c *Client) getURL(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", c.URL, path)
}

// POST 方法向指定路径发送POST请求
// 返回：状态码、响应内容、错误信息
func (c *Client) POST(path string, headers map[string]string, data io.Reader) (int, []byte, error) {
	req, err := http.NewRequest("POST", c.getURL(path), data)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return c.do(req)
}

// GET 方法向指定路径发送GET请求
// 返回：状态码、响应内容、错误信息
func (c *Client) GET(path string, headers map[string]string) (int, []byte, error) {
	req, err := http.NewRequest("GET", c.getURL(path), nil)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return c.do(req)

}

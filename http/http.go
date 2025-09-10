package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type HttpClient struct {
	baseURL string
	client  *http.Client
}

func NewHTTPClient(baseURL string) *HttpClient {
	return &HttpClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (h *HttpClient) DoAuthQuery(ctx context.Context, method, path, token string, query url.Values, resp interface{}) error {
	// 拼接 URL，如果有 query 参数
	fullURL := h.baseURL + path
	if len(query) > 0 {
		fullURL += "?" + query.Encode()
	}

	// 创建请求，body 为 nil
	req, err := http.NewRequestWithContext(ctx, method, fullURL, nil)
	if err != nil {
		return err
	}

	// 设置认证 Header
	req.Header.Set("OPEN-API-AUTH", token)

	// 发送请求
	return h.send(req, resp)
}

func (h *HttpClient) DoAuthForm(ctx context.Context, method, path, token string, data url.Values, resp interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, h.baseURL+path, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("OPEN-API-AUTH", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return h.send(req, resp)
}

func (h *HttpClient) DoAuthJSON(ctx context.Context, method, path, token string, data interface{}, resp interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, h.baseURL+path, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Set("OPEN-API-AUTH", token)
	req.Header.Set("Content-Type", "application/json")

	return h.send(req, resp)
}

// 统一执行逻辑
func (h *HttpClient) send(req *http.Request, resp interface{}) error {
	res, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(resp)
}

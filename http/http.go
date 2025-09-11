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

func (h *HttpClient) DoAuthQuery(ctx context.Context, method, path, token string, data interface{}, resp interface{}) error {
	values, err := h.structToValues(data)
	if err != nil {
		return err
	}

	// 拼接 URL，如果有 query 参数
	fullURL := h.baseURL + path
	if len(values) > 0 {
		fullURL += "?" + values.Encode()
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

func (h *HttpClient) DoAuthForm(ctx context.Context, method, path, token string, data interface{}, resp interface{}) error {
	values, err := h.structToValues(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, h.baseURL+path, strings.NewReader(values.Encode()))
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

func (h *HttpClient) structToValues(data interface{}) (url.Values, error) {
	// 先转成 map[string]interface{}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	// 转成 url.Values
	values := url.Values{}
	for k, v := range m {
		values.Add(k, fmt.Sprint(v))
	}
	return values, nil
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

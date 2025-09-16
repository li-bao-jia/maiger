package after

import (
	"context"
)

type SignRequest struct {
	AfterId string `json:"afterId"` // 售后ID
}

type SignResponse struct {
	Data    bool   `json:"data"`    // 返回内容 true 成功 false 失败
	Code    string `json:"code"`    // 返回编码
	Message string `json:"message"` // 返回说明
}

// 3.6 确认收货

func (a *AfterSaleService) Sign(ctx context.Context, token string, req *SignRequest) (resp *SignResponse, err error) {
	err = a.http.DoAuthQuery(ctx, "GET", "/open/api/after/sign", token, req, &resp)

	return resp, err
}

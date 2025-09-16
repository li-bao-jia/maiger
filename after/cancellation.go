package after

import (
	"context"
)

type CancellationRequest struct {
	AfterId string `json:"afterId"` // 售后ID
}

type CancellationResponse struct {
	Data    bool   `json:"data"`
	Code    string `json:"code"`    // 返回编码
	Message string `json:"message"` // 返回说明
}

// 3.5 取消售后

func (a *AfterSaleService) Cancellation(ctx context.Context, token string, req *CancellationRequest) (resp *CancellationResponse, err error) {
	err = a.http.DoAuthQuery(ctx, "GET", "/open/api/after/cancellation", token, req, &resp)

	return resp, err
}

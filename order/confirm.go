package order

import (
	"context"
)

type ConfirmRequest struct {
	OrderSn string `json:"orderSn"` // 订单号（提交订单返回的订单号）
}

type ConfirmResponse struct {
	Data    bool   `json:"data"`    // 状态
	Code    string `json:"code"`    // 返回编码
	Message string `json:"message"` // 返回说明
}

// 5.4 确认收货

// 订单商品状态为已签收后，可以通过确认收货接口修改订单已完成状态。
// 获取返回结果时 先判断HTTP请求状态码在判断业务返回编码
// 成功条件: HTTP请求状态码为200并且code为100时成功

func (o *OrderService) Confirm(ctx context.Context, token string, req *ConfirmRequest) (resp *ConfirmResponse, err error) {
	err = o.http.DoAuthQuery(ctx, "POST", "/open/api/order/confirmReceived", token, req, &resp)

	return resp, err
}

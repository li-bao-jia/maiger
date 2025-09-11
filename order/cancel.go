package order

import (
	"context"
)

type CancelRequest struct {
	ThirdOrder string `json:"thirdOrder"` // 第三⽅订单号（20个字符以内）
}

type CancelResponse struct {
	Data    bool   `json:"data"`
	Code    string `json:"code"`    // 返回编码
	Message string `json:"message"` // 返回说明
}

// 5.3 取消未支付订单

// 提交订单后，如果未支付取消订单，调用取消未支付订单接口进行订单取消操作。
// 取消订单后不做发货处理。已支付订单不能取消。如果未支付订单24小时未支付系统默认取消订单
// 获取返回结果时 先判断HTTP请求状态码在判断业务返回编码。
// 成功条件: HTTP请求状态码为200并且code为100时成功。

func (o *OrderService) Cancel(ctx context.Context, token string, req *CancelRequest) (resp *CancelResponse, err error) {
	err = o.http.DoAuthQuery(ctx, "POST", "/open/api/order/cancelOrder", token, req, &resp)

	return resp, err
}

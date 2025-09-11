package order

import (
	"context"
)

type PayRequest struct {
	ThirdOrder string `json:"thirdOrder"` // 第三⽅订单号（20个字符以内）
}

type PayResponse struct {
	Data    bool   `json:"data"`    // 状态
	Code    string `json:"code"`    // 返回编码
	Message string `json:"message"` // 返回说明
}

// 5.2 订单支付

// 提交订单成功后调用订单支付接口支付订单。支付扣取商城授信额度。额度不足的情况支付不成功，不支持发货。支付成功后才进行发货处理。
// 获取返回结果时 先判断HTTP请求状态码在判断业务返回编码。
// 成功条件: HTTP请求状态码为200并且code为100时成功。

func (o *OrderService) Pay(ctx context.Context, token string, req *PayRequest) (resp *PayResponse, err error) {
	err = o.http.DoAuthQuery(ctx, "POST", "/open/api/order/toPay", token, req, &resp)

	return resp, err
}

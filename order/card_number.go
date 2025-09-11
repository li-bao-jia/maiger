package order

import (
	"context"
)

type CardNumberRequest struct {
	OrderSn string `json:"orderSn"` // 订单号（提交订单返回的订单号）
}

type CardNumberResponse struct {
	Data    []CardNumberResult `json:"data"`
	Code    string             `json:"code"`
	Message string             `json:"message"`
}

type CardNumberResult struct {
	CardNumber     string `json:"cardNumber,omitempty"`     // 卡券或卡密
	CardCouponCode int    `json:"cardCouponCode,omitempty"` // 卡卷卷码
	FaceValue      string `json:"faceValue,omitempty"`      // 面值
	ExpireDate     string `json:"expireDate,omitempty"`     // 过期时间
}

// 5.9 查看虚拟订单卡券、卡密

// 提交订单支付成功后，虚拟商品订单，卡卷卡密类型商品可以通过该接口获取卡卷卡密信息。
// 获取返回结果时 先判断HTTP请求状态码在判断业务返回编码。
// 成功条件: HTTP请求状态码为200并且code为100时成功。

func (o *OrderService) CardNumber(ctx context.Context, token string, req *CardNumberRequest) (resp *CardNumberResponse, err error) {
	// 调用前给 resp 初始化
	// 防止传入的非虚拟卡券卡密单号，data 是 []（空数组）
	// JSON 解码器会尝试往 resp.Data 写入一个 slice，resp 本身是 nil，没有可写的结构体空间，直接报错
	resp = &CardNumberResponse{}

	err = o.http.DoAuthQuery(ctx, "GET", "/open/api/order/selectOrderCardNumber", token, req, resp)

	return
}

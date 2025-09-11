package order

import "context"

type StatusResponse struct {
	Data    []StatusResult `json:"data"`    // 状态
	Code    string         `json:"code"`    // 返回编码
	Message string         `json:"message"` // 返回说明
}

type StatusResult struct {
	OrderSn     string `json:"orderSn"`     // 订单号
	Freight     string `json:"freight"`     // 运费
	OrderAmount string `json:"orderAmount"` // 订单总额（包含运费)
	PayStatus   int    `json:"payStatus"`   // 支付状态：0-未付款；2-已付款
	OrderStatus int    `json:"orderStatus"` // 订单状态：订单状态(0-待付款；1-待发货；3-已发货 5-已取消；11-已完成)
}

// 5.8 订单状态查询

// 提交订单成功后，可以通过查询订单状态接口来判断订单是否已发货、已完成等订单状态信息订单。
// 获取返回结果时 先判断HTTP请求状态码在判断业务返回编码。
// 成功条件: HTTP请求状态码为200并且code为100时成功。
// 参数：订单号（提交订单返回的订单号）集合，最大支持20个订单

func (o *OrderService) Status(ctx context.Context, token string, req *[]string) (resp *StatusResponse, err error) {
	err = o.http.DoAuthJSON(ctx, "POST", "/open/api/order/selectOrderStatus", token, req, &resp)

	return resp, err
}

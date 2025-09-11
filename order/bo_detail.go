package order

import "context"

type BoDetailRequest struct {
	BlanketOrderSn string `json:"blanketOrderSn"` // 主订单号
}

type BoDetailResponse struct {
	Data    []Order `json:"data"`
	Code    string  `json:"code"`
	Message string  `json:"message"`
}

// 5.14 主订单查询详情

// 主订单查询详情接口，可以查看主订单下所有子订单订单状态以及订单商品信息以及发货状态信息。
// 获取返回结果时 先判断HTTP请求状态码在判断业务返回编码。
// 成功条件: HTTP请求状态码为200并且code为100时成功。

func (o *OrderService) BoDetail(ctx context.Context, token string, req *BoDetailRequest) (resp *BoDetailResponse, err error) {
	err = o.http.DoAuthQuery(ctx, "GET", "/open/api/order/selectBoDetail", token, req, &resp)

	return resp, err
}

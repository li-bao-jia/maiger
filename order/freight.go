package order

import (
	"context"
)

type FreightRequest struct {
	Address   string      `json:"address"`
	SkuIdList []SkuIdList `json:"skuIdList"`
}

type SkuIdList struct {
	SkuId    string `json:"skuId"`    // 商品SKU ID（只能是sku_id，spu_id提示商品不存在）
	Quantity int    `json:"quantity"` // 商品数量
}

type FreightResponse struct {
	Data    string `json:"data"`    // 运费
	Code    string `json:"code"`    // 返回编码
	Message string `json:"message"` // 返回说明
}

// 5.12 运费查询

// 通过运费查询接口，可以计算出当前提交订单商品所需要支付的运费金额。
// 获取返回结果时 先判断HTTP请求状态码在判断业务返回编码。
// 成功条件: HTTP请求状态码为200并且code为100时成功。

func (o *OrderService) Freight(ctx context.Context, token string, req *FreightRequest) (resp *FreightResponse, err error) {
	err = o.http.DoAuthJSON(ctx, "POST", "/open/api/order/productFreightQuery", token, req, &resp)

	return resp, err
}

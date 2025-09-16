package after

import (
	"context"
)

type QueryLogisticsRequest struct {
	OrderSn string `json:"orderSn"` // 订单编号
	SkuId   string `json:"skuId"`   // skuId
}

type QueryLogisticsResponse struct {
	Data    []QueryLogisticsResult `json:"data"`
	Code    string                 `json:"code"`    // 返回编码
	Message string                 `json:"message"` // 返回说明
}

type QueryLogisticsResult struct {
	SupplierID      string `json:"supplierId"`      // 供应商ID
	MaxSaleQuantity int    `json:"maxSaleQuantity"` // 最大可售后数量
	Quantity        int    `json:"quantity"`        // 申请数量
	LogisticsID     string `json:"logisticsId"`     // 物流ID
	LogisticsNo     string `json:"logisticsNo"`     // 物流号
	ShippingName    string `json:"shippingName"`    // 配送名称
	Status          int    `json:"status"`          // 状态：1: 无物流, 2: 在途, 3: 已签收
	StoreHouseID    string `json:"storeHouseId"`    // 发货仓库ID
}

// 3.8 根据订单商品ID获取发货包裹信息

func (a *AfterSaleService) QueryLogistics(ctx context.Context, token string, req *QueryLogisticsRequest) (resp *QueryLogisticsResponse, err error) {
	err = a.http.DoAuthQuery(ctx, "GET", "/open/api/after/queryLogistics", token, req, &resp)

	return resp, err
}

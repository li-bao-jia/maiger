package after

import (
	"context"
)

type AddLogisticsRequest struct {
	AfterId         string   `json:"afterId"`         // 售后ID
	SkuId           string   `json:"skuId"`           // 订单商品SkuId
	ShippingName    string   `json:"shippingName"`    // 配送名称
	LogisticsNo     string   `json:"logisticsNo"`     // 快递单号
	SendNumber      int      `json:"sendNumber"`      // 发货数量
	ShippingFeeType int16    `json:"shippingFeeType"` // 客户寄回快递支付情况
	ShippingFee     float64  `json:"shippingFee"`     // 运费
	SupplierId      string   `json:"supplierId"`      // 供应商ID
	StoreHouseId    string   `json:"storeHouseId"`    // 仓库ID
	LogisticsId     string   `json:"logisticsId"`     // 物流ID
	OrderGoodsId    string   `json:"orderGoodsId"`    // 订单商品ID
	PicUrls         []string `json:"picUrls"`         // 图片列表
}

type AddLogisticsResponse struct {
	Data    bool   `json:"data"`
	Code    string `json:"code"`    // 返回编码
	Message string `json:"message"` // 返回说明
}

// 3.7 添加物流信息

func (a *AfterSaleService) AddLogistics(ctx context.Context, token string, req *[]AddLogisticsRequest) (resp *AddLogisticsResponse, err error) {
	err = a.http.DoAuthJSON(ctx, "POST", "/open/api/after/addLogistics", token, req, &resp)

	return resp, err
}
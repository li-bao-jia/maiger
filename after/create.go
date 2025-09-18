package after

import (
	"context"
)

type CreateRequest struct {
	OrderSn      string         `json:"orderSn"`                // 订单编号
	SkuId        string         `json:"skuId"`                  // 商品sku
	ReturnType   int            `json:"returnType"`             // 售后要求 (0-退货/退款; 2-换货; 4-补寄; 7-退款)
	Reasons      string         `json:"reasons"`                // 申请原因
	QuestionDesc string         `json:"questionDesc,omitempty"` // 问题描述（非必填）
	PicList      []string       `json:"picList,omitempty"`      // 图片列表（非必填）
	SupplierList []SupplierList `json:"supplierList,omitempty"` // 供应商供货信息（非必填）
}

type SupplierList struct {
	SupplierId   string `json:"supplierId"`            // 供应商ID
	Quantity     int    `json:"quantity"`              // 申请数量
	StoreHouseId string `json:"storeHouseId"`          // 仓库ID
	LogisticsNo  string `json:"logisticsNo"`           // 快递单号
	LogisticsId  string `json:"logisticsId,omitempty"` // 快递单号id（非必填）
}

type CreateResponse struct {
	Data    []CreateResult `json:"data"`
	Code    string         `json:"code"`    // 返回编码
	Message string         `json:"message"` // 返回说明
}

type CreateResult struct {
	SkuId           string   `json:"skuId"`           // 商品skuId
	OrderGoodsId    string   `json:"orderGoodsId"`    // 订单商品ID
	AfterSaleIdList []string `json:"afterSaleIdList"` // 售后编号
}

// 3.2 创建售后单

func (a *AfterSaleService) Create(ctx context.Context, token string, req *[]CreateRequest) (resp *CreateResponse, err error) {
	err = a.http.DoAuthJSON(ctx, "POST", "/open/api/after/create", token, req, &resp)

	return resp, err
}

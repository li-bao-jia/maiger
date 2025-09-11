package order

import (
	"context"
)

type CheckAreaLimitRequest struct {
	Address string   `json:"address"`   // 收件地址
	SkuIds  []string `json:"skuIdList"` // 商品skuIds
}

type CheckAreaLimitResponse struct {
	Data    []CheckAreaLimitResult `json:"data"`
	Code    string                 `json:"code"`    // 返回编码
	Message string                 `json:"message"` // 返回说明
}

type CheckAreaLimitResult struct {
	SkuId          string `json:"skuId"`          // 商品 sku_Id
	IsAreaRestrict bool   `json:"isAreaRestrict"` // true 或空值代表区域受限 false 区域不受限
}

// 5.13 查询商品区域限制

func (o *OrderService) CheckAreaLimit(ctx context.Context, token string, req *CheckAreaLimitRequest) (resp *CheckAreaLimitResponse, err error) {
	err = o.http.DoAuthJSON(ctx, "POST", "/open/api/order/checkAreaLimit", token, req, &resp)

	return resp, err
}

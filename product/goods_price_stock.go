package product

import (
	"context"
)

type GoodsPriceStockRequest struct {
	SkuIds []string `json:"skuIdList"` // skuId数组
}

type GoodsPriceStockResponse struct {
	Data    []GoodsPriceStockResult `json:"data"`    // 分页数据
	Code    string                  `json:"code"`    // 返回编码
	Message string                  `json:"message"` // 返回说明
}

type GoodsPriceStockResult struct {
	GoodsNumber int    `json:"goodsNumber"` // 库存 非售卖为空
	MinNumber   int    `json:"minNumber"`   // 最小起购数量
	SkuId       string `json:"skuId"`       // 商品skuId
	SkuStatus   int    `json:"skuStatus"`   // 1:非售卖 2:上架
	SellPrice   string `json:"sellPrice"`   // 销售价 非售卖为空
}

// 4.7 批量查询价格库存

func (s *GoodsService) GoodsPriceStock(ctx context.Context, token string, req *GoodsPriceStockRequest) (resp *GoodsPriceStockResponse, err error) {
	err = s.http.DoAuthJSON(ctx, "POST", "/open/api/product/goods/batchPriceStock", token, req, &resp)

	return resp, err
}

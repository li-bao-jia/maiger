package product

import "context"

type SimilarRequest struct {
	SkuId string `json:"skuId"` // 商品skuId（必填）
}

type SimilarResponse struct {
	Data    []SimilarResult `json:"data"`    // 结果数据
	Code    string          `json:"code"`    // 返回编码
	Message string          `json:"message"` // 返回说明
}

type SimilarResult struct {
	SkuId     int    `json:"skuId"`     // 商品skuId
	SpecName  int    `json:"specName"`  // 规格名称
	SpecValue string `json:"specValue"` // 规格值
}

// 4.10 查询相似商品

func (s *GoodsService) Similar(ctx context.Context, token string, req *SimilarRequest) (resp *SimilarResponse, err error) {
	err = s.http.DoAuthQuery(ctx, "GET", "/open/api/product/goods/similarSku", token, req, &resp)

	return resp, err
}

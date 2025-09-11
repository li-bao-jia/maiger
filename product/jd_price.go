package product

import "context"

type JdPriceRequest struct {
	ThirdGoodsNo string `json:"thirdGoodsNo"` // 三方商品编号（必填）……4.4 查询商品详情 thirdId 字段
	GoodsUrl     string `json:"goodsUrl"`     // 商品链接（必填）……4.4 查询商品详情 thirdUrl 字段
}

type JdPriceResponse struct {
	Data    JdPriceResult `json:"data"`    // 结果数据
	Code    string        `json:"code"`    // 返回编码
	Message string        `json:"message"` // 返回说明
}

type JdPriceResult struct {
	JdPrice      int    `json:"jdPrice"`      // 京东商品价格
	JdSkuId      int    `json:"jdSkuId"`      // 京东skuId
	ThirdGoodsNo string `json:"thirdGoodsNo"` // 三方商品编号
}

// 4.9 查询京东商品价格

func (s *GoodsService) JdPrice(ctx context.Context, token string, req *[]JdPriceRequest) (resp *JdPriceResponse, err error) {
	err = s.http.DoAuthJSON(ctx, "POST", "/open/api/product/goods/getPrice", token, req, &resp)

	return resp, err
}

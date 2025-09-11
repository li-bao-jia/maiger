package product

import (
	"context"
)

type GoodsPageRequest struct {
	Lang      string `json:"lang"`      // 语言类型，默认中文，可选项：中文，english，french，portuguese
	PageNum   int    `json:"pageNum"`   // 页码
	PageSize  int    `json:"pageSize"`  // 每页数量
	GoodsName string `json:"goodsName"` // 商品名称
	BrandId   string `json:"brandId"`   // 品牌ID
	CatId     string `json:"catId"`     // 分类ID
}

type GoodsPageResponse struct {
	Data    GoodsPageResult `json:"data"`    // 分页数据
	Code    string          `json:"code"`    // 返回编码
	Message string          `json:"message"` // 返回说明
}

type GoodsPageResult struct {
	PageNum   int       `json:"pageNum"`   // 页码
	PageSize  int       `json:"pageSize"`  // 每页条数
	Total     int       `json:"total"`     // 总数
	TotalPage int       `json:"totalPage"` // 总页数
	Products  []Product `json:"data"`      // 商品列表
}

type Product struct {
	GoodsId         string `json:"goodsId"`         // 商品ID
	GoodsName       string `json:"goodsName"`       // 商品名称
	GoodsPrice      string `json:"goodsPrice"`      // 指导价
	MinNumber       int    `json:"minNumber"`       // 起订量
	SellPrice       string `json:"sellPrice"`       // 销售价
	SkuId           string `json:"skuId"`           // SKU ID
	ThumbnailImgUrl string `json:"thumbnailImgUrl"` // 主图
	ThirdPrice      string `json:"thirdPrice"`      // 第三方价格
}

// 4.3 查询商品列表(需注意，此接口按SKU纬度返回的数据)

func (s *GoodsService) GoodsPage(ctx context.Context, token string, req *GoodsPageRequest) (resp *GoodsPageResponse, err error) {
	err = s.http.DoAuthForm(ctx, "POST", "/open/api/product/goods/findGoodsPage", token, req, &resp)

	return resp, err
}

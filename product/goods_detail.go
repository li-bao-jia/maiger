package product

import (
	"context"
	"fmt"
	"net/url"
)

type GoodsDetailRequest struct {
	Lang    string `json:"lang"`    // 语言
	GoodsId string `json:"goodsId"` // 商品Id
	SkuId   string `json:"skuId"`   // 商品SKU ID(与商品ID必传一个)
}

type GoodsDetailResponse struct {
	Code    string            `json:"code"`    // 返回编码
	Message string            `json:"message"` // 具体内容
	Data    GoodsDetailResult `json:"data"`    // 返回数据
}

type GoodsDetailResult struct {
	BrandID                string             `json:"brandId"`                // 品牌Id
	ThirdBrandID           string             `json:"thirdBrandId"`           // 三方品牌id
	ThirdBrandName         string             `json:"thirdBrandName"`         // 三方品牌名称
	BrandName              string             `json:"brandName"`              // 品牌名称
	GoodsID                string             `json:"goodsId"`                // 商品Id
	GoodsName              string             `json:"goodsName"`              // 商品名称
	GoodsSn                string             `json:"goodsSn"`                // 货号
	IsReal                 int                `json:"isReal"`                 // 是否实物
	BarCode                string             `json:"barCode,omitempty"`      // 条形码（已作废）
	GoodsUnit              string             `json:"goodsUnit"`              // 商品单位
	CatID                  string             `json:"catId"`                  // 分类id
	TaxRate                string             `json:"taxRate"`                // 税率
	ThirdCatID             string             `json:"thirdCatId"`             // 三方三级分类id
	ThirdCatName           string             `json:"thirdCatName"`           // 三方三级分类名称
	ModelNumber            string             `json:"modelNumber"`            // 商品型号
	W3CCode                string             `json:"w3cCode"`                // W3C编码
	GoodsDescGalleryList   []GoodsDescGallery `json:"goodsDescGalleryList"`   // 商品详情图片
	GoodsGalleryList       []GoodsGallery     `json:"goodsGalleryList"`       // 商品图片列表
	GoodsSkuBaseInfoVoList []GoodsSkuBaseInfo `json:"goodsSkuBaseInfoVoList"` // SKU 基本信息
}

type GoodsGallery struct {
	ImgUrl        string `json:"imgUrl"`                  // 图片地址
	ImgSort       int    `json:"imgSort,omitempty"`       // 排序
	PicType       int    `json:"picType,omitempty"`       // 类型
	CoverImageUrl string `json:"coverImageUrl,omitempty"` // 封面图
	VideoUrl      string `json:"videoUrl,omitempty"`      // 视频 URL
}

type GoodsDescGallery struct {
	ImgUrl  string `json:"imgUrl"`            // 图片地址
	ImgSort int    `json:"imgSort,omitempty"` // 排序
}

type GoodsSkuBaseInfo struct {
	SkuID           string      `json:"skuId"`             // SKU ID
	SkuName         string      `json:"skuName"`           // SKU 名称
	BarCode         string      `json:"barCode,omitempty"` // 条形码
	ThumbnailImgUrl string      `json:"thumbnailImgUrl"`   // 主图
	EnableQuantity  int         `json:"enableQuantity"`    // 库存
	GoodsPrice      string      `json:"goodsPrice"`        // 指导价
	SellPrice       string      `json:"sellPrice"`         // 销售价
	CarriagePrice   string      `json:"carriagePrice"`     // 运费
	MinNumber       int         `json:"minNumber"`         // 起订量
	GoodsSpecList   []GoodsSpec `json:"goodsSpecList"`     // 商品规格信息
}

type GoodsSpec struct {
	SkuID     string `json:"skuId"`     // SKU ID
	SpecName  string `json:"specName"`  // 规格名称
	SpecValue string `json:"specValue"` // 规格值
}

// 4.4 查询商品详情

func (s *GoodsService) GoodsDetail(ctx context.Context, token string, req *GoodsDetailRequest) (resp *GoodsDetailResponse, err error) {
	p := map[string]interface{}{
		"lang":    req.Lang,
		"skuId":   req.SkuId,
		"goodsId": req.GoodsId,
	}

	values := url.Values{}
	for k, v := range p {
		values.Add(k, fmt.Sprint(v))
	}

	err = s.http.DoAuthQuery(ctx, "GET", "/open/api/product/goods/findDetail", token, values, &resp)

	return resp, err
}

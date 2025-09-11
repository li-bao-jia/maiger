package product

import (
	"context"
)

type GoodsDetailSpecRequest struct {
	GoodsId string `json:"goodsId"` // 商品Id
}

type GoodsDetailSpecResponse struct {
	Code    string                `json:"code"`    // 返回编码
	Message string                `json:"message"` // 具体内容
	Data    GoodsDetailSpecResult `json:"data"`    // 返回数据
}

type GoodsDetailSpecResult struct {
	GoodsId              string                 `json:"goodsId"`              // 商品Id
	IsReal               int                    `json:"isReal"`               // 是否真实商品
	RechargeModel        string                 `json:"rechargeModel"`        //
	VirtualSourceType    string                 `json:"virtualSourceType"`    //
	GoodsDesc            string                 `json:"goodsDesc"`            // 商品描述
	GoodsDescType        int                    `json:"goodsDescType"`        // 商品描述类型
	GoodsDescGalleryList []GoodsDescSpecGallery `json:"goodsDescGalleryList"` // 商品描述图片
}

type GoodsDescSpecGallery struct {
	ImgUrl  string `json:"imgUrl"`  // 图片地址
	ImgSort int    `json:"imgSort"` // 图片排序
}

// 4.5 查询商品描述

func (s *GoodsService) GoodsDetailSpec(ctx context.Context, token string, req *GoodsDetailSpecRequest) (resp *GoodsDetailSpecResponse, err error) {
	err = s.http.DoAuthQuery(ctx, "GET", "/open/api/product/goods/findDetailSpec", token, req, &resp)

	return resp, err
}

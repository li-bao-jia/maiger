package product

import (
	"context"
	"fmt"
	"net/url"
)

type BrandPageRequest struct {
	Lang      string `json:"lang"`      // 语言类型，默认中文，可选项：中文，english，french，portuguese
	PageNum   int    `json:"pageNum"`   // 页码
	PageSize  int    `json:"pageSize"`  // 每页数量
	BrandId   string `json:"brandId"`   // 品牌ID
	BrandName string `json:"brandName"` // 品牌名称
}

type BrandPageResponse struct {
	Data    BrandPageResult `json:"data"`    // 分页数据
	Code    string          `json:"code"`    // 返回编码
	Message string          `json:"message"` // 返回说明
}

type BrandPageResult struct {
	PageNum   int     `json:"pageNum"`   // 页码
	PageSize  int     `json:"pageSize"`  // 每页条数
	Total     int     `json:"total"`     // 总数
	TotalPage int     `json:"totalPage"` // 总页数
	Brands    []Brand `json:"data"`      // 品牌列表
}

type Brand struct {
	BrandId   string `json:"brandId"`   // 品牌ID
	BrandName string `json:"brandName"` // 品牌名称
}

// 4.1 查询品牌列表

func (s *GoodsService) BrandPage(ctx context.Context, token string, req *BrandPageRequest) (resp *BrandPageResponse, err error) {
	p := map[string]interface{}{
		"lang":      req.Lang,
		"pageNum":   req.PageNum,
		"pageSize":  req.PageSize,
		"brandId":   req.BrandId,
		"brandName": req.BrandName,
	}

	values := url.Values{}
	for k, v := range p {
		values.Add(k, fmt.Sprint(v))
	}

	err = s.http.DoAuthForm(ctx, "POST", "/open/api/product/brand/listPage", token, values, &resp)

	return resp, err
}

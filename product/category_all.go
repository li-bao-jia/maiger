package product

import (
	"context"
)

type Category struct {
	CatId    string     `json:"catId"`    // 分类ID
	CatName  string     `json:"catName"`  // 分类名称
	CatClass int        `json:"catClass"` // 分类名称
	ParentId string     `json:"parentId"` // 父级ID
	Children []Category `json:"children"` // 子级分类
}

type CategoryAllRequest struct {
	Lang string `json:"lang"` // 语言类型，默认中文，可选项：中文，english，french，portuguese
}

type CategoryAllResponse struct {
	Data    []Category `json:"data"`    // 分页数据
	Code    string     `json:"code"`    // 返回编码
	Message string     `json:"message"` // 返回说明
}

// 4.8 查询全量分类信息

func (s *GoodsService) CategoryAll(ctx context.Context, token string, req *CategoryAllRequest) (resp *CategoryAllResponse, err error) {
	err = s.http.DoAuthJSON(ctx, "GET", "/open/api/product/category/all", token, req, &resp)

	return resp, err
}

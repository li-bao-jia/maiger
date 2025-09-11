package product

import (
	"context"
)

type CategoryPageRequest struct {
	Lang     string `json:"lang"`     // 语言类型，默认中文，可选项：中文，english，french，portuguese
	PageNum  int    `json:"pageNum"`  // 页码
	PageSize int    `json:"pageSize"` // 每页数量
	ParentId string `json:"parentId"` // 父级分类ID
}

type CategoryPageResponse struct {
	Data    CategoryPageResult `json:"data"`    // 分页数据
	Code    string             `json:"code"`    // 返回编码
	Message string             `json:"message"` // 返回说明
}

type CategoryPageResult struct {
	PageNum    int        `json:"pageNum"`   // 页码
	PageSize   int        `json:"pageSize"`  // 每页条数
	Total      int        `json:"total"`     // 总数
	TotalPage  int        `json:"totalPage"` // 总页数
	Categories []Category `json:"data"`      // 分类列表
}

// 4.2 查询分类列表

func (s *GoodsService) CategoryPage(ctx context.Context, token string, req *CategoryPageRequest) (resp *CategoryPageResponse, err error) {
	err = s.http.DoAuthQuery(ctx, "GET", "/open/api/product/category/findByParentId", token, req, &resp)

	return resp, err
}

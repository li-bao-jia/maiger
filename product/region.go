package product

import (
	"context"
)

type RegionRequest struct {
	ParentId int `json:"parentId"` // 区域四级，第一级的父级id为1
}

type RegionResponse struct {
	Data    []RegionResult `json:"data"`    // 分页数据
	Code    string         `json:"code"`    // 返回编码
	Message string         `json:"message"` // 返回说明
}

type RegionResult struct {
	ParentId   int    `json:"parentId"`   // 父级id
	RegionId   int    `json:"regionId"`   // 区域id
	RegionName string `json:"regionName"` // 区域名称
	RegionType int    `json:"regionType"` // 区域层级
}

// 4.6 查询区域信息

func (s *GoodsService) Region(ctx context.Context, token string, req *RegionRequest) (resp *RegionResponse, err error) {
	err = s.http.DoAuthQuery(ctx, "GET", "/open/api/product/goods/region", token, req, &resp)

	return resp, err
}

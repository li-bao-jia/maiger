package after

import (
	"context"
)

type QueryApplyRequest struct {
	OrderSn string `json:"orderSn"` // 订单编号
	SkuId   string `json:"skuId"`   // skuId
}

type QueryApplyResponse struct {
	Data    []QueryApplyResult `json:"data"`
	Code    string             `json:"code"`    // 返回编码
	Message string             `json:"message"` // 返回说明
}

type QueryApplyResult struct {
	ReturnType     int    `json:"returnType"`     // 售后类型(0-退货/退款; 2-换货； 4-补寄； 6-补偿费用; 7-退款 )
	ReturnTypeName string `json:"returnTypeName"` // 售后类型名称
}

// 3.1 查询售后类型以及包裹信息

// 通过订单编号和订单商品ID获取可申请的售后类型,以及需要申请售后的包裹

func (a *AfterSaleService) QueryApply(ctx context.Context, token string, req *QueryApplyRequest) (resp *QueryApplyResponse, err error) {
	err = a.http.DoAuthQuery(ctx, "GET", "/open/api/after/queryApply", token, req, &resp)

	return resp, err
}

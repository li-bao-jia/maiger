package after

import (
	"context"
)

type CreateRefundRequest struct {
	OrderSn      string `json:"orderSn"`      // 订单编号
	Reasons      string `json:"reasons"`      // 售后原因
	ReturnType   int    `json:"returnType"`   // 售后类型：0退货退款，2换货，4补寄，6费用补偿，7退款
	QuestionDesc string `json:"questionDesc"` // 问题描述
}

type CreateRefundResponse struct {
	Data    CreateRefundResult `json:"data"`
	Code    string             `json:"code"`    // 返回编码
	Message string             `json:"message"` // 返回说明
}

type CreateRefundResult struct {
	AfterId string `json:"afterId"` // 售后id
	AfterSn string `json:"afterSn"` // 售后编号
}

// 3.9 创建退款售后(新版)

// 在待发货状态下申请退款，该退款是订单维度进行退款

func (a *AfterSaleService) CreateRefund(ctx context.Context, token string, req *CreateRefundRequest) (resp *CreateRefundResponse, err error) {
	err = a.http.DoAuthJSON(ctx, "POST", "/open/api/after/createRefund", token, req, &resp)

	return resp, err
}

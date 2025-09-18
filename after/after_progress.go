package after

import (
	"context"
)

type AfterProgressRequest struct {
	AfterId string `json:"afterId"`
}

type AfterProcessResponse struct {
	Code    string               `json:"code"`    // 返回编码
	Message string               `json:"message"` // 返回说明
	Data    []AfterProcessItemVo `json:"data"`    // 返回内容
}

type AfterProcessItemVo struct {
	AfterID          string `json:"afterId"`          // 售后ID
	ActionProcess    int    `json:"actionProcess"`    // 操作流程
	OrganizeName     string `json:"organizeName"`     // 组织名称
	Operator         string `json:"operator"`         // 操作人
	OperatorTime     string `json:"operatorTime"`     // 操作时间
	ActionResult     int    `json:"actionResult"`     // 审核结果
	AuditNode        bool   `json:"auditNode"`        // 是否审核节点
	SortNum          int    `json:"sortNum"`          // 排序
	ActionDesc       string `json:"actionDesc"`       // 审批意见/节点描述
	ActionName       string `json:"actionName"`       // 节点名称
	ReturnType       int    `json:"returnType"`       // 售后处理方式
	ProcessType      int    `json:"processType"`      // 售后节点类型
	LocusType        int    `json:"locusType"`        // 流程类型
	BusinessType     int    `json:"businessType"`     // 业务类型
	DeliveryType     string `json:"deliveryType"`     // 配送方式
	ReturnGoodsType  string `json:"returnGoodsType"`  // 退货类型
	OnlyRefundReason string `json:"onlyRefundReason"` // 仅退款原因
	LogisticsNo      string `json:"logisticsNo"`      // 物流单号
	ShippingName     string `json:"shippingName"`     // 快递公司
	WarehouseName    string `json:"warehouseName"`    // 仓库名称
	ShippingFee      string `json:"shippingFee"`      // 运费
}

// 3.10 查看进度（新版）

func (a *AfterSaleService) AfterProgress(ctx context.Context, token string, req *AfterProgressRequest) (resp *AfterProcessResponse, err error) {
	err = a.http.DoAuthQuery(ctx, "GET", "/open/api/after/afterProgress", token, req, &resp)
	return
}

package after

import "context"

type ViewProgressRequest struct {
	AfterId string `json:"afterId"` // 售后ID
}

type ViewProgressResponse struct {
	Code    string           `json:"code"`    // 返回编码
	Message string           `json:"message"` // 返回说明
	Data    []ViewProgressVo `json:"data"`    // 返回内容
}

type ViewProgressVo struct {
	ActionSysUsername string       `json:"actionSysuserName"` // 操作人名称
	OrganizeName      string       `json:"organizeName"`      // 组织名称
	ActionProcess     string       `json:"actionProcess"`     // 操作流程(1、提交售后2、项目审核售后 3、经销商审核售后 4、总站审核售后 5、购买人退件 6、供应商上签收 7、供应商发货 8、购买人签收 9、申请退款 10、审核退款)
	ActionName        string       `json:"actionName"`        // 操作名称
	ActionDesc        string       `json:"actionDesc"`        // 操作描述
	ActionResultDesc  string       `json:"actionResultDesc"`  // 操作结果说明
	OperationTime     string       `json:"operationTime"`     // 操作时间
	ClientType        int          `json:"clientType"`        // 操作端填写0
	SupplierSource    int          `json:"supplierSource"`    // 供应商来源 3：总站 2：经销商 1：项目
	ApplierPhone      string       `json:"applierPhone"`      // 客诉联系电话
	FastMailList      []FastMailVo `json:"fastMailList"`      // 物流信息
}

type FastMailVo struct {
	LogisticsId   string `json:"logisticsId"`   // 物流ID
	ShippingName  string `json:"shippingName"`  // 快递公司名称
	LogisticsNo   string `json:"logisticsNo"`   // 快递单号
	SendNumber    int    `json:"sendNumber"`    // 发货数量
	LogisticsType int    `json:"logisticsType"` // 单号类型 1、客户下单购买发货物流 2、客户寄回物流单号 3、上架补发货物流单号
	Status        int    `json:"status"`        // 物流状态 1、无物流 2、在途 3、已签收
}

// 3.3 查询进度（废弃）

func (a *AfterSaleService) ViewProgress(ctx context.Context, token string, req *SignRequest) (resp *ViewProgressResponse, err error) {
	err = a.http.DoAuthQuery(ctx, "GET", "/open/api/after/viewProgress", token, req, &resp)

	return resp, err
}

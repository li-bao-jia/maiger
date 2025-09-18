package order

import (
	"context"
)

type SubmitRequest struct {
	SkuList           []SkuItem `json:"skuList"`           // 下单商品信息 (必填)
	ThirdOrder        string    `json:"thirdOrder"`        // 第三方订单号 (必填)
	Name              string    `json:"name"`              // 收货人姓名 (必填)
	Mobile            string    `json:"mobile"`            // 电话号码 (必填)
	RechargeAccount   string    `json:"rechargeAccount"`   // 充值账号 (虚拟商品必填)
	Address           string    `json:"address"`           // 收货人详细地址 (非必填)
	ZipCode           string    `json:"zipCode"`           // 邮编 (非必填)
	City              string    `json:"city"`              // 收货人市级 (非必填)
	County            string    `json:"county"`            // 收货人县 (非必填)
	Email             string    `json:"email"`             // 邮箱 (非必填)
	Phone             string    `json:"phone"`             // 收货人座机号 (非必填)
	Province          string    `json:"province"`          // 收货人省份 (非必填)
	PurchaserName     string    `json:"purchaserName"`     // 采购单位 (非必填)
	PurchaserDeptName string    `json:"purchaserDeptName"` // 采购部门 (非必填)
	RealName          string    `json:"realName"`          // 下单人真实姓名 (非必填)
	UserAccount       string    `json:"userAccount"`       // 下单人账户 (非必填)
	UserPhone         string    `json:"userPhone"`         // 下单人手机号 (非必填)
	Auditor           string    `json:"auditor"`           // 审核人账户 (非必填)
	AuditContactName  string    `json:"auditContactName"`  // 审核联系人姓名 (非必填)
	AuditorPhone      string    `json:"auditorPhone"`      // 审核人手机号 (非必填)
	OrganizationName  string    `json:"organizationName"`  // 审核人机构名称 (非必填)
	Remark            string    `json:"remark"`            // 备注 (非必填)
	Town              string    `json:"town"`              // 收货人乡镇 (非必填)
}

type SkuItem struct {
	Num   int     `json:"num"`   // 商品数量 (必填)
	SkuId string  `json:"skuId"` // 商品SkuId (必填)
	Price float64 `json:"price"` // 商品价格 (非必填)
}

type SubmitResponse struct {
	Code    string         `json:"code"`    // 返回编码
	Message string         `json:"message"` // 返回说明
	Data    []SubmitResult `json:"data"`    // 订单内容
}

type SubmitResult struct {
	BlanketOrderSn  string                `json:"blanketOrderSn"`  // 提交单号返回
	MainOrderSn     string                `json:"mainOrderSn"`     // 主单号
	OrderID         string                `json:"orderId"`         // 订单ID
	OrderSn         string                `json:"orderSn"`         // 订单号
	Freight         string                `json:"freight"`         // 运费
	PayStatus       string                `json:"pay_status"`      // ?
	TotalOrderPrice string                `json:"totalOrderPrice"` // 订单总额（不含运费）
	SkuList         []SubmitResultSkuList `json:"skuList"`         // 订单对应商品信息
}

type SubmitResultSkuList struct {
	SkuID           string      `json:"skuId"`
	ClientSkuID     interface{} `json:"clientSkuId"`
	ThirdSubOrderID interface{} `json:"thirdSubOrderId"`
	Num             int         `json:"num"`
	Price           string      `json:"price"`
	GoodsRemark     interface{} `json:"goodsRemark"`
}

// 5.1 提交订单

// 提交订单接口。一次最多支持50个商品下单。
// 提交订单后根据系统会进行拆单处理，提交成功后会返回订单对应的拆单信息。需要第三方系统保存拆单的订单号，以便后续接口调用参数传值使用。
// 获取返回结果时 先判断HTTP请求状态码在判断业务返回编码
// 成功条件: HTTP请求状态码为200并且code为100时成功

func (o *OrderService) Submit(ctx context.Context, token string, req *SubmitRequest) (resp *SubmitResponse, err error) {
	err = o.http.DoAuthJSON(ctx, "POST", "/open/api/order/submitOrder", token, req, &resp)

	return resp, err
}

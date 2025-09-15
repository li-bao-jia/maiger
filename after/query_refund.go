package after

import "context"

type QueryRefundRequest struct {
	AfterId string `json:"afterId"` // 售后ID
}

type QueryRefundResponse struct {
	Data    QueryRefundResult `json:"data"`
	Code    string            `json:"code"`    // 返回编码
	Message string            `json:"message"` // 返回说明
}

type QueryRefundResult struct {
	AfterId           string   `json:"afterId"`           // 售后ID
	AfterSn           string   `json:"afterSn"`           // 售后编号
	MainOrderSn       string   `json:"mainOrderSn"`       // 主订单号
	OrderSn           string   `json:"orderSn"`           // 订单号
	AfterStatus       int      `json:"afterStatus"`       // 售后状态
	ReturnType        int      `json:"returnType"`        // 售后要求
	ApplierName       string   `json:"applierName"`       // 客诉联系人
	ApplierPhone      string   `json:"applierPhone"`      // 客诉联系电话
	FullAddress       string   `json:"fullAddress"`       // 全地址
	ThirdTradeNum     string   `json:"thirdTradeNum"`     // 第三方支付单号
	PaymentMode       string   `json:"paymentMode"`       // 支付模式（接口返回是字符串，保持 string）
	RefundType        int16    `json:"refundType"`        // 退款方式
	ShippingFee       string   `json:"shippingFee"`       // 扣除运费（返回是字符串，建议用 string 接收，必要时再转 decimal）
	MaxRefundCash     string   `json:"maxRefundCash"`     // 退款金额（返回是字符串）
	MaxRefundIntegral int      `json:"maxRefundIntegral"` // 退还积分
	ProjectName       string   `json:"projectName"`       // 项目名称
	UserAccount       string   `json:"userAccount"`       // 用户名称（可能为 null）
	Householder       string   `json:"householder"`       // 户主（可能为 null）
	OpeningBank       string   `json:"openingBank"`       // 开户行（可能为 null）
	RefundAccount     string   `json:"refundAccount"`     // 账号（可能为 null）
	CompensateReason  string   `json:"compensateReason"`  // 补偿原因
	PicUrls           []string `json:"picUrls"`           // 退款凭证图片列表
}

// 3.4 查看退款单

func (a *AfterSaleService) QueryRefund(ctx context.Context, token string, req *QueryRefundRequest) (resp *QueryRefundResponse, err error) {
	err = a.http.DoAuthQuery(ctx, "GET", "/open/api/after/queryRefund", token, req, &resp)

	return resp, err
}

package after

import "context"

type QueryTradeNoRequest struct {
	RefundFlowId string `json:"refundFlowId"` // 退款流水ID
}

type QueryTradeNoResponse struct {
	Code    string             `json:"code"`    // 返回编码
	Message string             `json:"message"` // 返回说明
	Data    QueryTradeNoResult `json:"data"`    // 返回内容
}

type QueryTradeNoResult struct {
	IIDD                  string                `json:"iidd"`                  // 售后id
	AfterSn               string                `json:"afterSn"`               // 售后编号
	Quantity              int                   `json:"quantity"`              // 售后总数量
	UserAccount           string                `json:"userAccount"`           // 用户账号
	Reasons               string                `json:"reasons"`               // 售后原因
	QuestionDesc          string                `json:"questionDesc"`          // 售后描述
	ReturnType            string                `json:"returnType"`            // 处理方式
	OldReturnType         string                `json:"oldReturnType"`         // 售后要求
	AfterStatus           int16                 `json:"afterStatus"`           // 售后状态
	CreateDate            string                `json:"createDate"`            // 创建时间
	UpdateDate            string                `json:"updateDate"`            // 修改时间
	RefundStatus          int16                 `json:"refundStatus"`          // 退款状态
	AuditStatus           int16                 `json:"auditStatus"`           // 审核状态
	Order                 AfterDetailOrderVo    `json:"order"`                 // 订单信息
	GoodsList             []AfterDetailGoodsVo  `json:"goodsList"`             // 售后商品列表
	OrderLogistics        AfterOrderLogisticsVo `json:"orderLogistics"`        // 订单物流信息
	UserLogisticsList     []AfterLogisticsVo    `json:"userLogisticsList"`     // 用户发货物流
	SupplierLogisticsList []AfterLogisticsVo    `json:"supplierLogisticsList"` // 供应商发货物流
	SupplierAddressVo     SupplierAddressVo     `json:"supplierAddressVo"`     // 供应商收货地址
}

type AfterDetailOrderVo struct {
	OrderSn       string `json:"orderSn"`       // 订单SN
	PayTime       string `json:"payTime"`       // 支付时间
	OrderStatus   int16  `json:"orderStatus"`   // 订单状态
	UserName      string `json:"userName"`      // 下单人姓名
	UserTelephone string `json:"userTelephone"` // 下单人手机号
	CreateDate    string `json:"createDate"`    // 创建时间
	UpdateDate    string `json:"updateDate"`    // 修改时间
	OrderID       string `json:"orderId"`       // 订单编号
}

type AfterDetailGoodsVo struct {
	GoodsName       string  `json:"goodsName"`       // 商品名称
	BrandName       string  `json:"brandName"`       // 品牌名称
	SkuID           string  `json:"skuId"`           // SKU ID
	GoodsSn         string  `json:"goodsSn"`         // 商品货号
	ModelNumber     string  `json:"modelNumber"`     // 商品型号
	GoodsAttr       string  `json:"goodsAttr"`       // 商品属性
	ThumbnailImgUrl string  `json:"thumbnailImgUrl"` // 商品缩略图
	Quantity        int     `json:"quantity"`        // 购买数量
	SaleQuantity    int     `json:"saleQuantity"`    // 售后数量
	Reality         int16   `json:"reality"`         // 是否实物
	ShippingStatus  int16   `json:"shippingStatus"`  // 配送状态
	RefundIntegral  int     `json:"refundIntergral"` // 退款积分
	RefundCash      float64 `json:"refundCash"`      // 退款现金
}

type AfterOrderLogisticsVo struct {
	LogisticsNo string `json:"logisticsNo"` // 物流号
	LogisticsID string `json:"logisticsId"` // 物流ID
}

type AfterLogisticsVo struct {
	ShippingName    string  `json:"shippingName"`    // 配送名称
	ShippingID      string  `json:"shippingId"`      // 配送方式ID
	DeliveryTime    string  `json:"deliveryTime"`    // 发货时间
	ShippingFee     float64 `json:"shippingFee"`     // 运费
	ShippingFeeType int16   `json:"shippingFeeType"` // 运费支付情况
	LogisticsNo     string  `json:"logisticsNo"`     // 物流号
	HasSign         bool    `json:"hasSign"`         // 是否签收
	SendNumber      int     `json:"sendNumber"`      // 发送数量
}

type SupplierAddressVo struct {
	UserName    string `json:"userName"`    // 收件人姓名
	UserPhone   string `json:"userPhone"`   // 收件人手机号
	Address     string `json:"address"`     // 地址（不含省市区）
	FullAddress string `json:"fullAddress"` // 详细地址
	CountryID   string `json:"countryId"`   // 国家ID
	Country     string `json:"country"`     // 国家
	ProvinceID  string `json:"provinceId"`  // 省份ID
	Province    string `json:"province"`    // 省份
	CityID      string `json:"cityId"`      // 城市ID
	City        string `json:"city"`        // 城市
	DistrictID  string `json:"districtId"`  // 地区ID
	District    string `json:"district"`    // 地区
	TownID      string `json:"townId"`      // 街道ID
	Town        string `json:"town"`        // 街道
}

// 3.12 根据退款流水查询售后详情（定制商城专用）

func (a *AfterSaleService) QueryTradeNo(ctx context.Context, token string, req *QueryTradeNoRequest) (resp *QueryTradeNoResponse, err error) {
	err = a.http.DoAuthJSON(ctx, "POST", "/open/api/after/queryTradeNo", token, req, &resp)

	return resp, err
}

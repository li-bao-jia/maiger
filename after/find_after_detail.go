package after

import (
	"context"
)

type FindAfterDetailRequest struct {
	AfterId string `json:"afterId"` // 售后ID
}

type FindAfterDetailResponse struct {
	Data    DetailResult `json:"data"`
	Code    string       `json:"code"`    // 返回编码
	Message string       `json:"message"` // 返回说明
}

type DetailResult struct {
	Iidd                  string                `json:"iidd"`                            // 售后id
	AfterSn               string                `json:"afterSn"`                         // 售后编号
	Quantity              int                   `json:"quantity"`                        // 售后总数量
	UserAccount           string                `json:"userAccount,omitempty"`           // 用户账号（非必填）
	Reasons               string                `json:"reasons"`                         // 售后原因
	QuestionDesc          string                `json:"questionDesc"`                    // 售后描述
	ReturnType            int                   `json:"returnType"`                      // 处理方式
	OldReturnType         int                   `json:"oldReturnType"`                   // 售后要求
	AfterStatus           int                   `json:"afterStatus"`                     // 售后状态 1-待处理；2-处理中；3-已处理；4-已完成；6-取消售后
	CreateDate            string                `json:"createDate"`                      // 创建时间
	UpdateDate            string                `json:"updateDate"`                      // 修改时间
	RefundStatus          int                   `json:"refundStatus,omitempty"`          // 退款状态（非必填）
	AuditStatus           int                   `json:"auditStatus,omitempty"`           // 审核状态（非必填）
	Order                 DetailOrder           `json:"order"`                           // 订单信息
	GoodsList             []DetailGoods         `json:"goodsList"`                       // 售后商品列表
	OrderLogistics        DetailLogistics       `json:"orderLogistics,omitempty"`        // 售后对应的订单的物流信息（非必填）
	UserLogisticsList     []DetailSendLogistics `json:"userLogisticsList,omitempty"`     // 用户发货的物流列表（非必填）
	SupplierLogisticsList []DetailSendLogistics `json:"supplierLogisticsList,omitempty"` // 供应商发货的物流列表（非必填）
	SupplierAddressVo     SupplierAddress       `json:"supplierAddressVo,omitempty"`     // 供应商收货地址（非必填）
}

type DetailOrder struct {
	OrderSn       string `json:"orderSn"`            // 订单sn
	PayTime       string `json:"payTime"`            // 订单支付时间
	OrderStatus   int    `json:"orderStatus"`        // 订单状态
	UserName      string `json:"userName,omitempty"` // 下单人姓名（非必填）
	UserTelephone string `json:"userTelephone"`      // 下单人手机号码
	CreateDate    string `json:"createDate"`         // 创建时间
	UpdateDate    string `json:"updateDate"`         // 修改时间
	OrderId       string `json:"orderId"`            // 订单编号
}

type DetailGoods struct {
	GoodsName       string `json:"goodsName"`              // 商品名称
	BrandName       string `json:"brandName"`              // 品牌名称
	SkuId           string `json:"skuId"`                  // SkuId
	GoodsSn         string `json:"goodsSn,omitempty"`      // 商品货号（非必填）
	ModelNumber     string `json:"modelNumber"`            // 商品型号
	GoodsAttr       string `json:"goodsAttr"`              // 购买该商品时所选择的属性
	ThumbnailImgUrl string `json:"thumbnailImgUrl"`        // 商品缩略图地址
	Quantity        int    `json:"quantity"`               // 购买数量
	SaleQuantity    int    `json:"saleQuantity,omitempty"` // 售后数量（非必填）
	Reality         int    `json:"reality"`                // 是否是实物，0：否，1：是
	ShippingStatus  int    `json:"shippingStatus"`         // 商品配送状态
}

type DetailLogistics struct {
	LogisticsNo string `json:"logisticsNo"` // 物流号
	LogisticsId string `json:"logisticsId"` // 物流单号ID
}

type DetailSendLogistics struct {
	ShippingName    string `json:"shippingName"`           // 配送名称
	ShippingId      string `json:"shippingId"`             // 配送方式ID
	DeliveryTime    string `json:"deliveryTime,omitempty"` // 发货时间（非必填）
	ShippingFee     string `json:"shippingFee"`            // 运费
	ShippingFeeType int    `json:"shippingFeeType"`        // 快递费支付情况（0-客户已付；1-到付）
	LogisticsNo     string `json:"logisticsNo"`            // 物流号
	HasSign         bool   `json:"hasSign"`                // 是否签收，true为签收，false为未签收
	SendNumber      int    `json:"sendNumber"`             // 发送数量
}

type SupplierAddress struct {
	UserName    string `json:"userName"`             // 收件人姓名
	UserPhone   string `json:"userPhone"`            // 收件人手机
	Address     string `json:"address"`              // 收件人地址（不包含省市区）
	FullAddress string `json:"fullAddress"`          // 收件人详细地址（包含省市区）
	CountryId   string `json:"countryId,omitempty"`  // 收货人的国家ID（非必填）
	Country     string `json:"country,omitempty"`    // 收货人的国家（非必填）
	ProvinceId  string `json:"provinceId,omitempty"` // 收货人的省份ID（非必填）
	Province    string `json:"province,omitempty"`   // 收货人的省份（非必填）
	CityId      string `json:"cityId,omitempty"`     // 收货人的城市ID（非必填）
	City        string `json:"city,omitempty"`       // 收货人的城市（非必填）
	DistrictId  string `json:"districtId,omitempty"` // 收货人的地区ID（非必填）
	District    string `json:"district,omitempty"`   // 收货人的地区（非必填）
	TownId      string `json:"townId,omitempty"`     // 收货人的街道/乡镇ID（非必填）
	Town        string `json:"town,omitempty"`       // 收货人的街道/乡镇（非必填）
}

// 3.11 查看售后详情（旧版）

func (a *AfterSaleService) FindAfterDetail(ctx context.Context, token string, req *FindAfterDetailRequest) (resp *FindAfterDetailResponse, err error) {
	err = a.http.DoAuthQuery(ctx, "GET", "/open/api/after/findAfterDetail", token, req, &resp)

	return resp, err
}

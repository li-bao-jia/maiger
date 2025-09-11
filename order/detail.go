package order

import (
	"context"
)

type DetailRequest struct {
	OrderSn string `json:"orderSn"` // 订单号（提交订单返回的订单号）
}

type DetailResponse struct {
	Data    Order  `json:"data"`    // 状态
	Code    string `json:"code"`    // 返回编码
	Message string `json:"message"` // 返回说明
}

type Order struct {
	OrderSn          string       `json:"orderSn"`          // 订单号
	Address          string       `json:"address"`          // 收货地址
	Consignee        string       `json:"consignee"`        // 收货人
	Phone            string       `json:"phone"`            // 收货人电话
	Freight          string       `json:"freight"`          // 运费
	GoodsAmount      string       `json:"goodsAmount"`      // 商品总额
	OrderAmount      string       `json:"orderAmount"`      // 订单支付现金
	OrderPoints      string       `json:"orderPoints"`      // 订单支付积分
	TotalOrderAmount string       `json:"totalOrderAmount"` // 订单总额（积分+现金）
	PayStatus        int          `json:"payStatus"`        // 支付状态：0-未付款；2-已付款
	OrderStatus      int          `json:"orderStatus"`      // 订单状态：订单状态(0-待付款；1-待发货；3-已发货 5-已取消；11-已完成)
	CreateDate       string       `json:"createDate"`       // 下单时间
	PayTime          string       `json:"payTime"`          // 支付时间（可选）
	ShippingTime     string       `json:"shippingTime"`     // 发货时间（可选）
	CompleteTime     string       `json:"completeTime"`     // 签收时间（可选）
	OrderGoodsList   []OrderGoods `json:"orderGoodsList"`   // 订单商品信息
}

type OrderGoods struct {
	GoodsName      string `json:"goodsName"`      // 商品名称
	BrandName      string `json:"brandName"`      // 商品品牌
	Freight        string `json:"freight"`        // 商品运费
	GoodsAmount    string `json:"goodsAmount"`    // 商品金额
	GoodsAttr      string `json:"goodsAttr"`      // 商品规格
	GoodsSn        string `json:"goodsSn"`        // 商品货号
	ModelNumber    string `json:"modelNumber"`    // 商品型号
	Quantity       int    `json:"quantity"`       // 商品数量
	ShippingStatus int    `json:"shippingStatus"` // 发货状态（0-未发货；1-部分发货；2-已发货；3-已签收）
	SkuId          string `json:"skuId"`          // 商品 SkuId
}

// 5.7 订单详情

func (o *OrderService) Detail(ctx context.Context, token string, req *DetailRequest) (resp *DetailResponse, err error) {
	err = o.http.DoAuthQuery(ctx, "GET", "/open/api/order/selectOrderDetail", token, req, &resp)

	return resp, err
}

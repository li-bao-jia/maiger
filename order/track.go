package order

import (
	"context"
)

type TrackRequest struct {
	OrderSn string `json:"orderSn"` // 订单号（提交订单返回的订单号）
}

type TrackResponse struct {
	Data    []TrackResult `json:"data"`
	Code    string        `json:"code"`    // 返回编码
	Message string        `json:"message"` // 返回说明
}

type TrackResult struct {
	LogisticsID         string         `json:"logisticsId"`        // 包裹ID
	DeliverType         int            `json:"deliverType"`        // 发货类型: 0:物流信息, 1:商家送货上门, 2:用户自取
	LogisticsNo         string         `json:"logisticsNo"`        // 运单号
	SkuList             []string       `json:"skuList"`            // 包裹关联的SKU列表
	Carrier             string         `json:"carrier"`            // 运输公司
	DeliveryPhone       string         `json:"deliveryPhone"`      // 送货人手机号码 (deliverType == 1 时有值)
	DeliveryCertificate string         `json:"deliveryCetificate"` // 送货凭证 (deliverType == 1 时有值)
	SignCertificate     string         `json:"signCetificate"`     // 签收凭证 (deliverType == 1 时有值)
	PackageCertificate  string         `json:"packageCetificate"`  // 打包凭证 (deliverType == 2 时有值)
	PickerPhone         string         `json:"pickerPhone"`        // 取货人手机号码 (deliverType == 2 时有值)
	PickCertificate     string         `json:"pickCetificate"`     // 取货凭证 (deliverType == 2 时有值)
	DeliveryPerson      string         `json:"deliveryPerson"`     // 送货人/取货人姓名 (根据 deliverType)
	Recipient           string         `json:"recipient"`          // 签收人姓名 (deliverType == 2 时有值)
	OrderTrackList      []DeliveryInfo `json:"orderTrackList"`     // 物流轨迹内容
}

// 5.6 查询订单配送信息

func (o *OrderService) Track(ctx context.Context, token string, req *TrackRequest) (resp *TrackResponse, err error) {
	err = o.http.DoAuthQuery(ctx, "GET", "/open/api/order/queryOrderTrack", token, req, &resp)

	return resp, err
}

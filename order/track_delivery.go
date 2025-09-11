package order

import (
	"context"
)

type TrackDeliveryRequest struct {
	Company string `json:"logisticsCompany"` // 物流公司(非必填，如果传递了，则会按照物流公司名称进行匹配搜索)
	Num     string `json:"logisticsNum"`     // 快递单号
}

type TrackDeliveryResponse struct {
	Data    TrackDeliveryResult `json:"data"`
	Code    string              `json:"code"`    // 返回编码
	Message string              `json:"message"` // 返回说明
}

type TrackDeliveryResult struct {
	DeliverType   int            `json:"deliverType"`    // 发货类型：0: 物流信息 1:商家送货上门 2:用户自取
	LogisticsId   string         `json:"logisticsId"`    // 包裹ID
	LogisticsNo   string         `json:"logisticsNo"`    // 运单号
	Carrier       string         `json:"carrier"`        // 运输公司
	SkuList       []string       `json:"skuList"`        // 包裹关联的SKU列表
	DeliveryInfos []DeliveryInfo `json:"orderTrackList"` // 物流轨迹内容
}

type DeliveryInfo struct {
	Time    string `json:"time"`    // 时间，原始格式
	FTime   string `json:"ftime"`   // 格式化后时间
	Status  string `json:"status"`  // 状态
	Context string `json:"context"` // 内容，物流情况
}

// 5.10 根据物流公司和物流单号查询配送信息

// 订单发货后，若无法调用查询订单配送信息接口，可调用该接口获取订单配送信息以及物流轨迹。
// 获取返回结果时 先判断HTTP请求状态码在判断业务返回编码。
// 成功条件: HTTP请求状态码为200并且code为100时成功。

func (o *OrderService) TrackDelivery(ctx context.Context, token string, req *TrackDeliveryRequest) (resp *TrackDeliveryResponse, err error) {
	err = o.http.DoAuthQuery(ctx, "POST", "/open/api/order/queryOrderTrackByCodeAndCompany", token, req, &resp)

	return resp, err
}

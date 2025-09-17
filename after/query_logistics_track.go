package after

import "context"

type QueryLogisticsTrackRequest struct {
	AfterId     string `json:"afterId"`     // 售后ID (与售后单号必须传一个)
	AfterSn     string `json:"afterSn"`     // 售后单号
	LogisticsNo string `json:"logisticsNo"` // 物流单号（必填）
}

type QueryLogisticsTrackResponse struct {
	Code    string                    `json:"code"`    // 返回编码
	Message string                    `json:"message"` // 返回说明
	Data    QueryLogisticsTrackResult `json:"data"`    // 返回内容
}

type QueryLogisticsTrackResult struct {
	LogisticsId    string             `json:"logisticsId"`    // 物流包裹ID
	Carrier        string             `json:"carrier"`        // 运输公司
	LogisticsNo    string             `json:"logisticsNo"`    // 运单号
	OrderTrackList []LogisticsTrackVo `json:"orderTrackList"` // 物流轨迹列表
}

type LogisticsTrackVo struct {
	Time    string `json:"time"`    // 原始时间
	FTime   string `json:"ftime"`   // 格式化时间
	Context string `json:"context"` // 物流情况
	Status  string `json:"status"`  // 状态
}

// 3.14 查询物流轨迹（新版本）

func (a *AfterSaleService) QueryLogisticsTrack(ctx context.Context, token string, req *QueryLogisticsTrackRequest) (resp *QueryLogisticsTrackResponse, err error) {
	err = a.http.DoAuthJSON(ctx, "POST", "/open/api/after/queryLogisticsTrack", token, req, &resp)

	return resp, err
}

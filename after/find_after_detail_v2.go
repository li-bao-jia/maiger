package after

import "context"

type FindAfterDetailV2Request struct {
	AfterId string `json:"afterId"` // 售后ID（售后ID与售后单号必须传一个）
	AfterSn string `json:"afterSn"` // 售后单号
}

// 3.13 查看售后详情（新版本）

func (a *AfterSaleService) FindAfterDetailV2(ctx context.Context, token string, req *FindAfterDetailV2Request) (resp *FindAfterDetailResponse, err error) {
	err = a.http.DoAuthJSON(ctx, "POST", "/open/api/after/v2/findAfterDetail", token, req, &resp)

	return resp, err
}

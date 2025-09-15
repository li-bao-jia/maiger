package after

import (
	"context"

	"github.com/li-bao-jia/maiger/http"
)

var (
	domain      = "https://www.xxxxxxxx.com"
	accessToken = "your_token"

	ctx              = context.Background()
	httpClient       = http.NewHTTPClient(domain)
	afterSaleService = NewAfterSaleService(httpClient)
)

// 测试 3.1 查询售后类型以及包裹信息

func TestQueryApply(t *testing.T) {
	queryApplyResp, err := afterSaleService.QueryApply(ctx, accessToken, &QueryApplyRequest{
		OrderSn: "25090715411241962663",
		SkuId:   "93da7c58-ecca-4fa9-901e-5b68bdfe1697",
	})

	if err != nil {
		t.Fatalf("QueryApply error: %v", err)
	}
	fmt.Println("QueryApply:", queryApplyResp.Data)
}

// 测试 3.4 查看退款单

func TestQueryRefund(t *testing.T) {
	queryRefundResp, err := afterSaleService.QueryRefund(ctx, accessToken, &QueryRefundRequest{
		AfterId: "375c7506-82fc-45ee-9d70-eea773f072e9",
	})

	if err != nil {
		t.Fatalf("QueryRefund error: %v", err)
	}
	fmt.Println("QueryRefund:", queryRefundResp.Data)
}

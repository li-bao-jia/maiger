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

// 测试 3.5 取消售后

func TestCancellation(t *testing.T) {
	cancelResp, err := afterSaleService.Cancellation(ctx, accessToken, &CancellationRequest{
		AfterId: "0a0a011a-8add-473f-8156-5157d5d48144",
	})

	if err != nil {
		t.Fatalf("Cancel error: %v", err)
	}
	fmt.Println("Cancel:", cancelResp.Data)
}

// 测试 3.6 确认收货

func TestSign(t *testing.T) {
	signResp, err := afterSaleService.Sign(ctx, accessToken, &SignRequest{
		AfterId: "0a0a011a-8add-473f-8156-5157d5d4814411",
	})

	if err != nil {
		t.Fatalf("Sign error: %v", err)
	}
	fmt.Println("Sign:", signResp.Data)
}

// 测试 3.8 根据订单商品ID获取发货包裹信息

func TestQueryLogistics(t *testing.T) {
	queryLogisticsResp, err := afterSaleService.QueryLogistics(ctx, accessToken, &QueryLogisticsRequest{
		OrderSn: "25090715411241962663",
		SkuId:   "93da7c58-ecca-4fa9-901e-5b68bdfe1697",
	})

	if err != nil {
		t.Fatalf("QueryLogistics error: %v", err)
	}
	fmt.Println("QueryLogistics:", queryLogisticsResp.Data)
}

// 测试 3.11 查看售后详情（旧版）

func TestFindAfterDetail(t *testing.T) {
	findAfterDetailResp, err := afterSaleService.FindAfterDetail(ctx, accessToken, &FindAfterDetailRequest{
		AfterId: "a1c50eae-62d7-4ddc-a208-83f024732386",
	})

	if err != nil {
		t.Fatalf("FindAfterDetail error: %v", err)
	}
	fmt.Println("FindAfterDetail:", findAfterDetailResp.Data)
}

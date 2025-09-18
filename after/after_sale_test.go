package after

import (
	"context"
	"fmt"
	"testing"

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

// 测试 3.2 创建售后单

func TestCreate(t *testing.T) {
	createResp, err := afterSaleService.Create(ctx, accessToken, &[]CreateRequest{
		{
			OrderSn:      "25090715411241962663",
			SkuId:        "93da7c58-ecca-4fa9-901e-5b68bdfe1697",
			ReturnType:   1,
			Reasons:      "商品缺少",
			QuestionDesc: "商品缺少",
			PicList:      []string{},
			SupplierList: []SupplierList{}, //创建退款的时候supplierList不用填写，其他申请必须填写
		},
	})

	if err != nil {
		t.Fatalf("Create error: %v", err)
	}
	fmt.Println("Create:", createResp.Data)
}

// 测试 3.3 查询进度（废弃）

func TestViewProgress(t *testing.T) {
	viewProgressResp, err := afterSaleService.ViewProgress(ctx, accessToken, &SignRequest{
		AfterId: "ec4d992d-0b1a-486d-8d02-1701c19fba7c",
	})

	if err != nil {
		t.Fatalf("ViewProgress error: %v", err)
	}
	fmt.Println("ViewProgress:", viewProgressResp.Data)
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

// 测试 3.10 查看进度（新版）

func TestAfterProgress(t *testing.T) {
	afterProgressResp, err := afterSaleService.AfterProgress(ctx, accessToken, &AfterProgressRequest{
		AfterId: "a1c50eae-62d7-4ddc-a208-83f024732386",
	})

	if err != nil {
		t.Fatalf("AfterProgress error: %v", err)
	}
	fmt.Println("AfterProgress:", afterProgressResp.Data)
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

// 测试 3.13 查看售后详情（新版本）

func TestFindAfterDetailV2(t *testing.T) {
	findAfterDetailResp, err := afterSaleService.FindAfterDetailV2(ctx, accessToken, &FindAfterDetailV2Request{
		AfterId: "a1c50eae-62d7-4ddc-a208-83f024732386",
		AfterSn: "",
	})

	if err != nil {
		t.Fatalf("FindAfterDetail error: %v", err)
	}
	fmt.Println("FindAfterDetail:", findAfterDetailResp.Data)
}

// 测试 3.14 查询物流轨迹（新版本）

func TestQueryLogisticsTrack(t *testing.T) {
	queryLogisticsTrackResp, err := afterSaleService.QueryLogisticsTrack(ctx, accessToken, &QueryLogisticsTrackRequest{
		AfterId:     "ec4d992d-0b1a-486d-8d02-1701c19fba7c",
		AfterSn:     "",
		LogisticsNo: "YT2562935022692",
	})

	if err != nil {
		t.Fatalf("QueryLogisticsTrack error: %v", err)
	}
	fmt.Println("QueryLogisticsTrack:", queryLogisticsTrackResp.Data)
}

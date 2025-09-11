package order

import (
	"context"
	"fmt"
	"testing"

	"github.com/li-bao-jia/maiger/http"
)

var (
	domain      = "https://www.xxxxxxxx.com"
	accessToken = "your_token"

	ctx          = context.Background()
	httpClient   = http.NewHTTPClient(domain)
	orderService = NewOrderService(httpClient)
)

// 测试 5.4 确认收货
func TestConfirm(t *testing.T) {
	confirmResp, err := orderService.Confirm(ctx, accessToken, &ConfirmRequest{
		OrderSn: "25082918411684343910",
	})

	if err != nil {
		t.Fatalf("Confirm error: %v", err)
	}
	fmt.Println("Confirm:", confirmResp.Data)
}

// 测试 5.5 批量确认收货

func TestConfirmBatch(t *testing.T) {
	confirmBatchResp, err := orderService.ConfirmBatch(ctx, accessToken, &[]string{
		"25082918411684343910",
	})

	if err != nil {
		t.Fatalf("ConfirmBatch error: %v", err)
	}
	fmt.Println("ConfirmBatch:", confirmBatchResp.Data)
}

// 测试 5.6 查询订单配送信息

func TestTrack(t *testing.T) {
	trackResp, err := orderService.Track(ctx, accessToken, &TrackRequest{
		OrderSn: "SN202509100000001",
	})

	if err != nil {
		t.Fatalf("Track error: %v", err)
	}
	fmt.Println("Track:", trackResp.Data)
}

// 测试 5.7 订单详情

func TestDetail(t *testing.T) {
	detailResp, err := orderService.Detail(ctx, accessToken, &DetailRequest{
		OrderSn: "25082918411684343910",
	})

	if err != nil {
		t.Fatalf("Detail error: %v", err)
	}
	fmt.Println("Detail:", detailResp.Data)
}

// 测试 5.8 订单状态查询

func TestStatus(t *testing.T) {
	statusResp, err := orderService.Status(ctx, accessToken, &[]string{
		"25082918411684343910",
	})

	if err != nil {
		t.Fatalf("Status error: %v", err)
	}
	fmt.Println("Status:", statusResp.Data)
}

// 测试 5.9 查看虚拟订单卡券、卡密

func TestCardNumber(t *testing.T) {
	cardNumberResp, err := orderService.CardNumber(ctx, accessToken, &CardNumberRequest{
		OrderSn: "25082918411684343910",
	})

	if err != nil {
		t.Fatalf("CardNumber error: %v", err)
	}
	fmt.Println("CardNumber:", cardNumberResp.Data)
}

// 测试 5.10 根据物流公司和物流单号查询配送信息

func TestTrackDelivery(t *testing.T) {
	trackDeliveryResp, err := orderService.TrackDelivery(ctx, accessToken, &TrackDeliveryRequest{
		Company: "顺丰",
		Num:     "SF202509100000001",
	})

	if err != nil {
		t.Fatalf("TrackDelivery error: %v", err)
	}
	fmt.Println("TrackDelivery:", trackDeliveryResp.Data)
}

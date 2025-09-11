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

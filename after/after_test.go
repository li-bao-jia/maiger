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

func TestApply(t *testing.T) {
	applyResp, err := afterSaleService.Apply(ctx, accessToken, &ApplyRequest{
		OrderSn: "25090715411241962663",
		SkuId:   "93da7c58-ecca-4fa9-901e-5b68bdfe1697",
	})

	if err != nil {
		t.Fatalf("Apply error: %v", err)
	}
	fmt.Println("Apply:", applyResp.Data)
}

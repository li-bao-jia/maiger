package coupon

import (
	"context"

	"github.com/li-bao-jia/maiger/http"
)

var (
	domain      = "https://www.xxxxxxxx.com"
	accessToken = "your_token"

	ctx           = context.Background()
	httpClient    = http.NewHTTPClient(domain)
	couponService = NewCouponService(httpClient)
)

package coupon

import "github.com/li-bao-jia/maiger/http"

type CouponService struct {
	http *http.HttpClient
}

func NewCouponService(http *http.HttpClient) *CouponService {
	return &CouponService{http: http}
}

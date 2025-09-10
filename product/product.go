package product

import "github.com/li-bao-jia/maiger/http"

type GoodsService struct {
	http *http.HttpClient
}

func NewGoodsService(http *http.HttpClient) *GoodsService {
	return &GoodsService{http: http}
}

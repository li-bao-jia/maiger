package order

import "github.com/li-bao-jia/maiger/http"

type OrderService struct {
	http *http.HttpClient
}

func NewOrderService(http *http.HttpClient) *OrderService {
	return &OrderService{http: http}
}

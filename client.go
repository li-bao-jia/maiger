package maiger

import (
	"github.com/li-bao-jia/maiger/after"
	"github.com/li-bao-jia/maiger/http"
	"github.com/li-bao-jia/maiger/message"
	"github.com/li-bao-jia/maiger/order"
	"github.com/li-bao-jia/maiger/product"
	"github.com/li-bao-jia/maiger/token"
)

type Client struct {
	AfterSale *after.AfterSaleService
	Message   *message.MessageService
	Order     *order.OrderService
	Product   *product.GoodsService
	Token     *token.TokenService
}

func NewClient(baseURL string) *Client {
	httpClient := http.NewHTTPClient(baseURL)

	return &Client{
		AfterSale: after.NewAfterSaleService(httpClient),
		Message:   message.NewMessageService(httpClient),
		Order:     order.NewOrderService(httpClient),
		Product:   product.NewGoodsService(httpClient),
		Token:     token.NewTokenService(httpClient),
	}
}

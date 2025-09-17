package after

import "github.com/li-bao-jia/maiger/http"

type AfterSaleService struct {
	http *http.HttpClient
}

func NewAfterSaleService(http *http.HttpClient) *AfterSaleService {
	return &AfterSaleService{http: http}
}

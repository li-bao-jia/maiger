package activity

import "github.com/li-bao-jia/maiger/http"

type ActivityService struct {
	http *http.HttpClient
}

func NewActivityService(http *http.HttpClient) *ActivityService {
	return &ActivityService{http: http}
}

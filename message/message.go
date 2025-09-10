package message

import "github.com/li-bao-jia/maiger/http"

type MessageService struct {
	http *http.HttpClient
}

func NewMessageService(http *http.HttpClient) *MessageService {
	return &MessageService{http: http}
}

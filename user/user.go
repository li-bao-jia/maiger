package user

import "github.com/li-bao-jia/maiger/http"

type UserService struct {
	http *http.HttpClient
}

func NewUserService(http *http.HttpClient) *UserService {
	return &UserService{http: http}
}

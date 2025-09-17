package token

import (
	"context"

	"github.com/li-bao-jia/maiger/http"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type TokenService struct {
	http *http.HttpClient
}

func NewTokenService(http *http.HttpClient) *TokenService {
	return &TokenService{http: http}
}

func (t *TokenService) GetToken(ctx context.Context, clientId, clientSecret string) (resp *TokenResponse, err error) {
	params := map[string]interface{}{
		"grant_type":    "client_credentials",
		"client_id":     clientId,
		"client_secret": clientSecret,
	}

	err = t.http.DoForm(ctx, "POST", "/open/api/oauth/token", params, &resp)

	return resp, err
}

func (t *TokenService) RefreshToken(ctx context.Context, clientId, clientSecret, refreshToken string) (resp *TokenResponse, err error) {
	params := map[string]interface{}{
		"grant_type":    "refresh_token",
		"client_id":     clientId,
		"client_secret": clientSecret,
		"refresh_token": refreshToken,
	}

	err = t.http.DoForm(ctx, "POST", "/open/api/oauth/token", params, &resp)

	return resp, err
}

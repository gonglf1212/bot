package botsdk

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bot/internal/bot/token"
	"github.com/go-resty/resty/v2"
)

const (
	// 获得token
	gatewayURI  string = "https://api.sgroup.qq.com/gateway"
	messagesURI string = "https://api.sgroup.qq.com/channels/{channel_id}/messages"
	// 用户注销
	contentType  string = "application/x-www-form-urlencoded"
	acceptHeader string = "application/json"
)

type Client struct {
	token       *token.Token
	gatewayURI  string
	revokeURL   string
	header      map[string]string
	client      *http.Client
	restyClient *resty.Client
}

func NewBotSdk(tokenParam *token.Token) *Client {
	client := &Client{
		token:      tokenParam,
		gatewayURI: gatewayURI,
		header: map[string]string{
			"content-type": contentType,
			"accept":       acceptHeader,
		},
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	authToken := fmt.Sprintf("%v.%s", tokenParam.AppID, tokenParam.AccessToken)
	client.restyClient = resty.New().
		SetTimeout(3 * time.Second).
		SetAuthToken(authToken).
		SetAuthScheme(string(tokenParam.Type))
	return client
}

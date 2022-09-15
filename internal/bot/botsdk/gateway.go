/*
 * @Author: gonglf
 * @Date: 2022-09-15 09:35:25
 * @LastEditors: gonglf
 * @LastEditTime: 2022-09-15 10:29:29
 * @Description:
 *
 */
package botsdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bot/dto"
	"github.com/bot/internal/bot/token"
	"github.com/go-resty/resty/v2"
)

const (
	// 获得token
	gatewayURI string = "https://api.sgroup.qq.com/gateway"
	// 用户注销
	contentType  string = "application/x-www-form-urlencoded"
	acceptHeader string = "application/json"
)

type client struct {
	token       *token.Token
	gatewayURI  string
	revokeURL   string
	header      map[string]string
	client      *http.Client
	restyClient *resty.Client
}

func NewBotSdk(tokenParam *token.Token) *client {
	client := &client{
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

func (c *client) Gateway() (*dto.WebsocketAP, error) {
	resp, err := c.restyClient.R().SetResult(dto.WebsocketAP{}).Get(c.gatewayURI)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Response Info:")
	// fmt.Println("  Error      :", err)
	// fmt.Println("  Status Code:", resp.StatusCode())
	// fmt.Println("  Status     :", resp.Status())
	// fmt.Println("  Body       :\n", resp)
	return resp.Result().(*dto.WebsocketAP), nil
}

func (c *client) doRequest(url string, data url.Values) ([]byte, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	for k, v := range c.header {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintln("StatusCode is ", res.StatusCode))
	}
	rebyte, err := ioutil.ReadAll(res.Body)
	fmt.Println("----111---", string(rebyte))
	return rebyte, err
}

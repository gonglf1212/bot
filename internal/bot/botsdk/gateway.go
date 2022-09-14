package botsdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// 获得token
	gatewayURI string = "https://api.sgroup.qq.com/gateway"
	// 用户注销
	contentType  string = "application/x-www-form-urlencoded"
	acceptHeader string = "application/json"
)

type client struct {
	gatewayURI string
	revokeURL  string
	header     map[string]string
	client     *http.Client
}

func NewBotSdk() *client {
	client := &client{
		gatewayURI: gatewayURI,
		header: map[string]string{
			"content-type": contentType,
			"accept":       acceptHeader,
		},
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	return client
}

func (c *client) Gateway(Authorization string) {
	if Authorization == "" {
		fmt.Println("-------")
		return
	}
	c.header["Authorization"] = Authorization
	c.doRequest(gatewayURI, nil)

}

func (c *client) doRequest(url string, data url.Values) ([]byte, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	fmt.Println("---333s----")

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

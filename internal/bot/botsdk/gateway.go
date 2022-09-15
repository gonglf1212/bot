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
	"github.com/bot/dto"
)

func (c *Client) Gateway() (*dto.WebsocketAP, error) {
	resp, err := c.restyClient.R().SetResult(dto.WebsocketAP{}).Get(c.gatewayURI)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.WebsocketAP), nil
}

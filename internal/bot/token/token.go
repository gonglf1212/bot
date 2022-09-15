/*
 * @Author: gonglf
 * @Date: 2022-09-15 09:35:25
 * @LastEditors: gonglf
 * @LastEditTime: 2022-09-15 10:28:38
 * @Description:
 *
 */
package token

import "fmt"

// Type token 类型
type Type string

// TokenType
const (
	TypeBot    Type = "Bot"
	TypeNormal Type = "Bearer"
)

type Token struct {
	AppID       uint64 `yaml:"appid"`
	AccessToken string `yaml:"token"`
	Type        Type   `yaml:"type"`
}

func (t *Token) GetString() string {
	if t.Type == TypeNormal {
		return t.AccessToken
	}
	return fmt.Sprintf("%v.%s", t.AppID, t.AccessToken)
}

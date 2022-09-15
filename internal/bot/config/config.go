/*
 * @Author: gonglf
 * @Date: 2022-09-15 09:35:25
 * @LastEditors: gonglf
 * @LastEditTime: 2022-09-15 10:30:33
 * @Description:
 *
 */
package config

import (
	"github.com/bot/internal/bot/token"
)

//全局访问
var (
	ConfigFile string
	BotToken   *token.Token
)

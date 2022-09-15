/*
 * @Author: gonglf
 * @Date: 2022-09-15 09:35:25
 * @LastEditors: gonglf
 * @LastEditTime: 2022-09-15 10:32:28
 * @Description:
 *
 */
package bot

import (
	"io/ioutil"

	"github.com/bot/internal/bot/config"
	"github.com/bot/internal/bot/token"
	"gopkg.in/yaml.v2"
)

//初始化
func Init() (err error) {

	bs, err := ioutil.ReadFile(config.ConfigFile)
	if err != nil {
		panic(err)
	}
	token := new(token.Token)
	err = yaml.Unmarshal(bs, token)
	if err == nil {
		config.BotToken = token
	}
	return
}

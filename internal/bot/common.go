/*
 * @Author: gonglf
 * @Date: 2022-09-15 09:35:25
 * @LastEditors: gonglf
 * @LastEditTime: 2022-09-15 17:15:54
 * @Description:
 *
 */
package bot

import (
	"io/ioutil"

	"github.com/bot/internal/bot/botsdk"
	"github.com/bot/internal/bot/config"
	"github.com/bot/internal/bot/token"
	"github.com/gocpp/log"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

//初始化
func Init() (err error) {

	if err = loadConfig(); err != nil {
		log.Error("loadConfig err", zap.Error(err))
		panic("loadConfig err")
	}

	if err = wsInfo(); err != nil {
		log.Error("wsInfo err", zap.Error(err))
		panic("wsInfo err")
	}

	return
}

func loadConfig() (err error) {

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

func wsInfo() (err error) {
	wsInfo, err := botsdk.NewBotSdk(config.BotToken).Gateway()
	log.Info("wsInfo", zap.Any("wsinfo", wsInfo.URL), zap.Any("err", err))
	config.WsInfo = wsInfo
	return
}

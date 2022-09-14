package bot

import (
	"io/ioutil"

	"github.com/bot/internal/core/config"
	"gopkg.in/yaml.v2"
)

//初始化
func Init() (err error) {

	bs, err := ioutil.ReadFile(config.ConfigFile)
	if err != nil {
		panic(err)
	}
	token := new(config.Token)
	err = yaml.Unmarshal(bs, token)
	if err == nil {
		config.BotToken = token
	}
	return
}

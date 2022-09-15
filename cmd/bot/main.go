/*
 * @Author: gonglf
 * @Date: 2022-09-15 09:35:25
 * @LastEditors: gonglf
 * @LastEditTime: 2022-09-15 14:21:54
 * @Description:
 *
 */
package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bot/internal/bot"
	"github.com/bot/internal/bot/config"
	"github.com/gocpp/log"
	"go.uber.org/zap"
)

func init() {
	flag.StringVar(&config.ConfigFile, "configfile", "./config/config.yaml", "config file")
}

func main() {
	bot.Init()

	server := new(bot.Server)
	go server.Connect(config.WsInfo, config.BotToken)

	handlerSignal()
}

func handlerSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGUSR1)

	for {
		s := <-c
		log.Info("Server get a signal ", zap.String("sig", s.String()))
		switch s {
		case syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//服务退出
			return
		case syscall.SIGUSR1:
			log.Info("Server config loading...")
		case syscall.SIGHUP:
			log.Info("Server restart...")
		default:
			return
		}
	}
}

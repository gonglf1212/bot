package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gocpp/log"
	"go.uber.org/zap"
)

func init() {

}

func main() {

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

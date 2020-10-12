package safeexit

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Capture(beforeExit func()) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	log.Println("应用已启动")

	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
			log.Println("收到退出信号:", s)
			beforeExit()
			return
		}
	}
}

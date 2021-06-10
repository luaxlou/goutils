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

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				log.Println(s,"waiting for exit..." )
				beforeExit()
				log.Println(s,"exited." )

				os.Exit(0)
			}
		}
	}()

}

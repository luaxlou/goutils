package main

import (
	"log"

	"github.com/luaxlou/goutils/safeexit"
)

func  main() {

	safeexit.Capture(func() {
		log.Println("before exit")
	})
}

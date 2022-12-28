package main

import (
	"github.com/mnes/haas/src/services/imaging-study/config"
)

func main() {
	config.SetEnv()
	setServer()
}

func setServer() {
	s := newServer()
	s.setConfig()
	s.register()
	go s.start()
	s.notifySignal()
	s.gracefulStop()
}

package main

import (
	"github.com/glebateee/core/logging"
	"github.com/glebateee/core/services"
)

func writeMessage(logger logging.Logger) {
	logger.Info("SportsStore")
}

func main() {
	services.RegisterDefaultServices()
	services.Call(writeMessage)
}

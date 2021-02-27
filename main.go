package main

import (
	"bitcoin/config"
	"bitcoin/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}

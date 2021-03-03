package main

import (
	"bitcoin/backend/controllers"
	"bitcoin/config"
	"bitcoin/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}

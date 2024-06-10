package main

import (
	"github.com/gauravsingh4242/bakar/config"
	"github.com/gauravsingh4242/bakar/controller"
	"github.com/gauravsingh4242/bakar/logger"
	"github.com/gauravsingh4242/bakar/routes"
)

func main() {
	logger.InitLogger("debug")
	defer logger.Sync()

	err := config.InitConfig()
	if err != nil {
		logger.Log.Debugf("error loading config: %s", err.Error())
		panic(err.Error())
	}

	bakarController := controller.NewBakarController()
	engine := routes.InitRoutes(bakarController)
	err = engine.Run(":8000")
	if err != nil {
		logger.Log.Errorf("Unable to start gin server: %s", err.Error())
		panic(err.Error())
	}

}

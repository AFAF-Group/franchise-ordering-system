package main

import (
	"fmt"

	"afaf-group.com/pkg/config"
	"afaf-group.com/pkg/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	app := echo.New()
	config.LoadEnv("config.env")
	conf := config.NewConfig()

	_, errMainDB := config.DBConnection(&conf.MainDatabase)
	if errMainDB != nil {
		log.Fatalf("Error connection to main db %v\n", errMainDB)
	}

	middlewares.LoggerMiddlewares(app)

	// run server
	address := fmt.Sprintf(":%s", conf.AppPort)

	if err := app.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}

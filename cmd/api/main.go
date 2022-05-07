package main

import (
	"fmt"

	"afaf-group.com/pkg/config"
	"afaf-group.com/pkg/middlewares"
	"afaf-group.com/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	app := echo.New()
	config.LoadEnv("config.env")
	conf := config.NewConfig()

	db, errMainDB := config.DBConnection(&conf.MainDatabase)
	if errMainDB != nil {
		log.Fatalf("Error connection to main db %v\n", errMainDB)
	}

	middlewares.LoggerMiddlewares(app)

	// Custom Validator
	// validator := config.NewCustomValidator()
	// app.Validator = validator
	// app.Use(func(handle echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(ctx echo.Context) error {
	// 		ctx.Set("validator", validator)
	// 		return handle(ctx)
	// 	}
	// })

	// Auth Routes
	router.InitAuthRoutes(app, db)

	// run server
	address := fmt.Sprintf(":%s", conf.AppPort)

	if err := app.Start(address); err != nil {
		log.Info("shutting down the server")
	}

}

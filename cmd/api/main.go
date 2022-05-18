package main

import (
	"fmt"

	"afaf-group.com/domain"
	"afaf-group.com/pkg/config"
	"afaf-group.com/pkg/middlewares"
	"afaf-group.com/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "afaf-group.com/docs/swagger" // docs is generated by Swag CLI, you have to import it.
)

// @title           Swagger Franchise Ordering System API
// @version         1.0
// @description     This is a franchise-ordering-system server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      franchise-ordering-system.swagger.io
// @BasePath  /v2
func main() {
	app := echo.New()
	config.LoadEnv("config.env")
	conf := config.NewConfig()

	db, errMainDB := config.DBConnection(&conf.MainDatabase)
	if errMainDB != nil {
		log.Fatalf("Error connection to main db %v\n", errMainDB)
	}

	app.HTTPErrorHandler = middlewares.NewHttpErrorHandler(domain.NewErrorStatusCodeMaps()).Handler
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
	// init Swagger
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	// run server
	address := fmt.Sprintf(":%s", conf.AppPort)

	if err := app.Start(address); err != nil {
		log.Info("shutting down the server")
	}

}

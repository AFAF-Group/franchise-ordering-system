package main

import (
	"fmt"
	"net/http"
	"strings"

	stdLog "log"

	"afaf-group.com/domain"
	"afaf-group.com/pkg/config"
	"afaf-group.com/pkg/middlewares"
	"afaf-group.com/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"

	_ "afaf-group.com/docs/swagger" // docs is generated by Swag CLI, you have to import it.
	_ "afaf-group.com/domain/models"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	// Instance Echo
	app := echo.New()

	// Load Env
	config.LoadEnv("config.env")

	conf := config.NewConfig()
	conf.Logger = config.ZapLoggerConfig(app, conf)
	logger := conf.Logger

	db, errMainDB := config.DBConnection(&conf.MainDatabase)
	if errMainDB != nil {
		log.Fatalf("Error connection to main db %v\n", errMainDB)
	}

	app.HTTPErrorHandler = middlewares.NewHttpErrorHandler(domain.NewErrorStatusCodeMaps()).Handler
	// Logger
	app.Use(middlewares.LoggerContextMiddleware(logger))
	app.Use(middlewares.LoggerMiddleware(&middlewares.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		RequestBodySkipper: func(httpRequest *http.Request) bool {
			if strings.HasPrefix(httpRequest.URL.Path, "/auth/login") {
				return true
			}
			return middlewares.DefaultLoggerRequestBodySkipper(httpRequest)
		},
		ResponseBodySkipper: func(httpRequest *http.Request, httpResponse *echo.Response) bool {
			if strings.HasPrefix(httpRequest.URL.Path, "/auth/login") {
				return true
			}
			return middlewares.DefaultLoggerResponseBodySkipper(httpRequest, httpResponse)
		},
	}))

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

	// Food Routes
	router.InitFoodRoutes(app, db)

	// Order Routes
	router.InitOrderRoutes(app, db)

	// Customer Routes
	router.InitCustomerRoutes(app, db)

	// init Swagger
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	// run server
	address := fmt.Sprintf(":%s", conf.AppPort)

	if err := app.Start(address); err != nil {
		log.Info("shutting down the server")
	}

	// Close DB
	defer closeDB(db, "Close DB Connection")

}

func closeDB(db *gorm.DB, message string) {
	stdLog.Println(message)

	sqlDB, err := db.DB()
	if err != nil {
		stdLog.Printf("Error close db connection %v\n", err)
	}

	if err := sqlDB.Close(); err != nil {
		stdLog.Printf("Error close db connection %v", err)
	}
}

package main

import (
	"afaf-group.com/pkg/config"
	"github.com/labstack/echo/v4"

	stdLog "log"
)

func main() {
	app := echo.New()
	config.LoadEnv(".env")
	conf := config.NewConfig()

	_, errMainDB := config.DBConnection(&conf.MainDatabase)
	if errMainDB != nil {
		stdLog.Fatalf("Error connection to main db %v\n", errMainDB)
	}

	app.Use()
}

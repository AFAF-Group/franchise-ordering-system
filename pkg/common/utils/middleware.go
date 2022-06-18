package utils

import (
	"afaf-group.com/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthMidlleware() echo.MiddlewareFunc {
	return middleware.JWT([]byte(domain.SECRET_JWT))
}

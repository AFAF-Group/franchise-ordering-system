package middlewares

import (
	"context"

	"afaf-group.com/pkg/common/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func LoggerContextMiddleware(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, constants.CTXKeyLogger, log)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

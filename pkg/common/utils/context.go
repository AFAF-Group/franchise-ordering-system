package utils

import (
	"context"

	"afaf-group.com/pkg/common/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func LoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(constants.CTXKeyLogger).(*zap.Logger)
	if ok {
		return logger
	}
	return nil
}

func LoggerFromEchoContext(ctx echo.Context) *zap.Logger {
	logger, ok := ctx.Request().Context().Value(constants.CTXKeyLogger).(*zap.Logger)
	if ok {
		return logger
	}
	return nil
}

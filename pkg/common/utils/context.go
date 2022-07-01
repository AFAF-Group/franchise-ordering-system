package utils

import (
	"context"

	"afaf-group.com/pkg/common/constants"
	"go.uber.org/zap"
)

func LoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(constants.CTXKeyLogger).(*zap.Logger)
	if ok {
		return logger
	}
	return nil
}

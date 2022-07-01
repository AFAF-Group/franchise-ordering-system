package config

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func ZapLoggerConfig(app *echo.Echo, conf *Config) *zap.Logger {
	var zapEncoder = zapLoggerEncoder()
	var zapWriteSyncer = zapLogWriter(conf)

	var logLevel = zapcore.InfoLevel
	if app.Debug {
		logLevel = zapcore.DebugLevel
	}

	var zapCore = zapcore.NewCore(zapEncoder, zapWriteSyncer, logLevel)
	var logger = zap.New(zapCore, zap.AddCaller())

	defer zapLoggerSync(logger)

	zap.ReplaceGlobals(logger)

	return logger
}

func zapLoggerSync(logger *zap.Logger) {
	if err := logger.Sync(); err != nil {
		log.Printf("error sync zap logger %v", err)
	}
}

func zapLoggerEncoder() zapcore.Encoder {
	var encoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func zapLogWriter(conf *Config) zapcore.WriteSyncer {
	var filename = fmt.Sprintf("%s/%s", conf.LogPath, conf.LogFile)
	return zapcore.AddSync(
		io.MultiWriter(
			zapcore.Lock(os.Stdout),
			&lumberjack.Logger{
				Filename:   filename,
				MaxSize:    50, // megabytes
				MaxBackups: 3,
				MaxAge:     28, // days
				LocalTime:  true,
			},
		),
	)
}

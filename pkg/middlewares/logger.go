package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"afaf-group.com/pkg/common/utils"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	LoggerConfig struct {
		Skipper             middleware.Skipper
		RequestBodySkipper  LoggerRequestBodySkipper
		ResponseBodySkipper LoggerResponseBodySkipper
		BodyReplacer        LoggerBodyReplacer
	}
	bodyDumpResponseWriter struct {
		io.Writer
		http.ResponseWriter
	}
	LoggerRequestBodySkipper  func(req *http.Request) bool
	LoggerResponseBodySkipper func(req *http.Request, res *echo.Response) bool
	LoggerBodyReplacer        func(body []byte) []byte
)

func DefaultLoggerRequestBodySkipper(req *http.Request) bool {
	contentType := req.Header.Get(echo.HeaderContentType)
	switch {
	case strings.HasPrefix(contentType, echo.MIMEApplicationForm),
		strings.HasPrefix(contentType, echo.MIMEMultipartForm):
		return true
	default:
		return false
	}
}

func DefaultLoggerResponseBodySkipper(req *http.Request, res *echo.Response) bool {
	return false
}

func DefaultLoggerBodyReplacer(body []byte) []byte {
	if body != nil {
		var responseBody map[string]interface{}
		if err := utils.JSONUnmarshal(body, &responseBody); err != nil {
			zap.L().Error("Error unmarshal response body", zap.Error(err))
		}

		if responseBody["data"] != nil {
			data, ok := responseBody["data"].(map[string]interface{})
			if ok {
				for k := range data {
					if k == "access_token" {
						data["access_token"] = "******"
					}
				}
				responseBody["data"] = data
			}
		}
		var sanitizedBody, err = utils.JSONMarshal(responseBody)
		if err != nil {
			zap.L().Error("Error marshal response body", zap.Error(err))
		}
		return sanitizedBody
	}
	return body
}

func LoggerMiddleware(config *LoggerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log := utils.LoggerFromContext(c.Request().Context())
			if config.Skipper == nil {
				config.Skipper = middleware.DefaultSkipper
			}
			if config.RequestBodySkipper == nil {
				config.RequestBodySkipper = DefaultLoggerRequestBodySkipper
			}
			if config.ResponseBodySkipper == nil {
				config.ResponseBodySkipper = DefaultLoggerResponseBodySkipper
			}
			if config.BodyReplacer == nil {
				config.BodyReplacer = DefaultLoggerBodyReplacer
			}
			start := time.Now()

			fields := []zapcore.Field{
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", time.Since(start).String()),
				zap.String("host", c.Request().Host),
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().RequestURI),
				zap.Int("status", c.Response().Status),
				zap.Int64("size", c.Response().Size),
				zap.String("user_agent", c.Request().UserAgent()),
			}

			id := c.Request().Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = c.Response().Header().Get(echo.HeaderXRequestID)
			}
			fields = append(fields, zap.String("request_id", id))

			if !config.RequestBodySkipper(c.Request()) {
				var reqBody []byte
				if c.Request().Body != nil { // Read
					reqBody, _ = ioutil.ReadAll(c.Request().Body)
					c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset
					if len(reqBody) > 0 {
						if resty.IsJSONType(c.Request().Header.Get(echo.HeaderContentType)) {
							fields = append(fields, zap.Any("request_body", json.RawMessage(reqBody)))
						}
					}
				}
			}

			// Response
			resBody := new(bytes.Buffer)
			if !config.ResponseBodySkipper(c.Request(), c.Response()) {
				mw := io.MultiWriter(c.Response().Writer, resBody)
				writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
				c.Response().Writer = writer
			}

			if err = next(c); err != nil {
				c.Error(err)
			}

			if len(resBody.Bytes()) > 0 {
				if resty.IsJSONType(c.Response().Header().Get(echo.HeaderContentType)) {
					fields = append(fields, zap.Any("response_body", json.RawMessage(config.BodyReplacer(resBody.Bytes()))))
				}
			}

			n := c.Response().Status
			switch {
			case n >= 500:
				log.With(zap.Error(err)).Error("Server error", fields...)
			case n >= 400:
				log.With(zap.Error(err)).Warn("Client error", fields...)
			case n >= 300:
				log.Info("Redirection", fields...)
			default:
				log.Info("Success", fields...)
			}

			return
		}
	}
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

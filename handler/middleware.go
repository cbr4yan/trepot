package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info("request",
				zap.String("uri", v.URI),
				zap.String("method", v.Method),
				zap.Int("status", v.Status),
				zap.String("remote_ip", v.RemoteIP),
			)
			return nil
		},
	})
}

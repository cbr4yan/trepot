package handler

import (
	"net/http"

	"github.com/cbr4yan/trepot/config"
	"github.com/cbr4yan/trepot/handler/company"
	"github.com/cbr4yan/trepot/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var (
	log = zap.L().Named("handler")
)

func New(cfg *config.Config, serviceProvider *service.Provider) *Api {
	return &Api{
		serviceProvider: serviceProvider,
	}
}

type Api struct {
	serviceProvider *service.Provider
}

func (a *Api) Handler() http.Handler {
	r := echo.New()

	r.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info("request", zap.String("uri", v.URI), zap.Int("status", v.Status))
			return nil
		},
	}))

	r.GET("/__health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	v1 := r.Group("/v1")
	{
		company.Register(v1, a.serviceProvider.CompanyService)
	}

	return r
}

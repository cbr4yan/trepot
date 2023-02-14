package http

import (
	"context"
	"net/http"
	"time"

	"github.com/cbr4yan/trepot/config"
	"go.uber.org/zap"
)

var (
	log = zap.L().Named("http")
)

const (
	defaultShutdownTimeout = 30 * time.Second
)

func New(cfg *config.Config, handler http.Handler) *server {
	return &server{
		httpServer: &http.Server{
			Addr:    cfg.HTTPServer.Addr,
			Handler: handler,
		},
	}
}

type server struct {
	httpServer *http.Server
}

func (s *server) Start() error {
	log.Info("starting http server", zap.String("addr", s.httpServer.Addr))
	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *server) Stop() error {
	log.Info("stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}

package command

import (
	"context"

	"github.com/cbr4yan/trepot/adapter/database"
	"github.com/cbr4yan/trepot/adapter/logger"
	"github.com/cbr4yan/trepot/adapter/server/http"
	"github.com/cbr4yan/trepot/config"
	"github.com/cbr4yan/trepot/handler"
	"github.com/cbr4yan/trepot/repo"
	"github.com/cbr4yan/trepot/service"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	serveCmd = cobra.Command{
		Use:  "serve",
		Long: "start API server",
		Run: func(cmd *cobra.Command, args []string) {
			runServe(cmd.Context())
		},
	}
)

func runServe(ctx context.Context) {
	cfg, err := config.Load("BACKEND")
	if err != nil {
		log.Fatal("failed to load config", zap.Error(err))
	}

	if err := logger.Setup(cfg); err != nil {
		log.Fatal("failed to setup logger", zap.Error(err))
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal("failed to setup database", zap.Error(err))
	}
	repoProvider := repo.New(db)
	serviceProvider := service.New(repoProvider)
	router := handler.New(cfg, serviceProvider)
	server := http.New(cfg, router.Handler())

	go func() {
		<-ctx.Done()
		if err := server.Stop(); err != nil {
			log.Fatal("failed to stop server", zap.Error(err))
		}
	}()

	if err := server.Start(); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}

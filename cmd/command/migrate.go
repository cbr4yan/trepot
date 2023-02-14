package command

import (
	"github.com/cbr4yan/trepot/adapter/database"
	"github.com/cbr4yan/trepot/adapter/logger"
	"github.com/cbr4yan/trepot/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	migrateCmd = cobra.Command{
		Use:  "migrate",
		Long: "migrate database structure",
		Run:  runMigrate,
	}
)

func runMigrate(cmd *cobra.Command, args []string) {
	cfg, err := config.Load("BACKEND")
	if err != nil {
		log.Fatal("failed to load config", zap.Error(err))
	}

	if err := logger.Setup(cfg); err != nil {
		log.Fatal("failed to setup logger", zap.Error(err))
	}

	engine, err := database.New(cfg)
	if err != nil {
		log.Fatal("failed to setup database", zap.Error(err))
	}
	defer engine.Close()

	driver, err := postgres.WithInstance(engine.DB, &postgres.Config{})
	if err != nil {
		log.Fatal("failed to setup migrate driver", zap.Error(err))
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("failed to setup migrate", zap.Error(err))
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to migrate", zap.Error(err))
	}

	log.Info("migrate success")
}

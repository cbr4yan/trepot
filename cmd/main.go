package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/cbr4yan/trepot/cmd/command"
	"go.uber.org/zap"
)

var (
	log = zap.L().Named("main")
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	defer cancel()

	go func() {
		<-ctx.Done()
		log.Info("gracefully stopping... (press Ctrl+C again to force)")
	}()

	if err := command.Register().ExecuteContext(ctx); err != nil {
		log.Fatal("failed to execute command", zap.Error(err))
	}
}

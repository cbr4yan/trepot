package command

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	log = zap.L().Named("command")

	rootCmd = cobra.Command{
		Use:     "backend",
		Version: "0.0.1",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func Register() *cobra.Command {
	rootCmd.AddCommand(&migrateCmd, &serveCmd)
	return &rootCmd
}

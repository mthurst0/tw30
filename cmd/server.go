package cmd

import (
	"github.com/mthurst0/tw30/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func serverCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "server",
		Short: "Run tw30 as a service",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := server.Run(); err != nil {
				zap.L().Fatal("failed to run service", zap.Error(err))
			}
			return nil
		},
	}
	return c
}

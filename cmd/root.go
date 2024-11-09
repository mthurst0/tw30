package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tw30",
		Short: "Toaster Whatever is 30",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd.Help()
			return nil
		},
		SilenceUsage: true,
	}

	rootCmd.AddCommand(serverCmd())
	rootCmd.AddCommand(versionCmd())

	return rootCmd
}

func Execute() {
	err := Root().Execute()
	if err != nil {
		os.Exit(1)
	}
}

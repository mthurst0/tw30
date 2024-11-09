package cmd

import (
	"github.com/mthurst0/tw30/internal/appenv"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Version command",
		RunE: func(cmd *cobra.Command, args []string) error {
			v := appenv.GetVersionInfo()
			cmd.Printf("Version    : %s\n", v.Version)
			cmd.Printf("Commit     : %s\n", v.Commit)
			cmd.Printf("Build time : %s\n", v.BuildTime)
			cmd.Printf("Go version : %s\n", v.GoVersion)
			return nil
		},
	}
}

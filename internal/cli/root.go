package cli

import (
	"context"

	"github.com/spf13/cobra"
)

// NewRootCmd creates the root command for the dive CLI.
func NewRootCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dive",
		Short: "A fast, minimal Docker/database CLI",
		Long:  `dive is a developer CLI for fast access to Docker containers and database CLIs.`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newPsCmd(ctx),
		newPgCmd(ctx),
		newRedisCmd(ctx),
		newMysqlCmd(ctx),
		newShellCmd(ctx),
		newLogsCmd(ctx),
	)

	return cmd
}

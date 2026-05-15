package cli

import (
	"context"

	"github.com/mzulfanw/dive/pkg/output"

	"github.com/mzulfanw/dive/internal/resolver"

	"github.com/mzulfanw/dive/internal/docker"

	"github.com/spf13/cobra"
)

// newLogsCmd creates the 'logs' command for streaming container logs.
func newLogsCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "logs <container>",
		Short:   "Stream logs from a container",
		Example: "dive logs api",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			container, err := resolver.ResolveShell(ctx, args[0])
			if err != nil {
				return output.FriendlyError("Could not find a matching container.", err)
			}
			return docker.StreamLogs(ctx, container)
		},
	}
	return cmd
}

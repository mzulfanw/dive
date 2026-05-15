package cli

import (
	"context"

	"github.com/mzulfanw/dive/internal/docker"
	"github.com/mzulfanw/dive/internal/resolver"
	"github.com/mzulfanw/dive/pkg/output"

	"github.com/spf13/cobra"
)

// newShellCmd creates the 'sh' command for shell access.
func newShellCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sh <container>",
		Short:   "Enter shell in a container (bash/sh)",
		Example: "dive sh api",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			container, err := resolver.ResolveShell(ctx, args[0])
			if err != nil {
				return output.FriendlyError("Could not find a matching container.", err)
			}
			return docker.ExecShell(ctx, container)
		},
	}
	return cmd
}

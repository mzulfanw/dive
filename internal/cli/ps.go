package cli

import (
	"context"

	"github.com/mzulfanw/dive/internal/docker"
	"github.com/mzulfanw/dive/pkg/output"

	"github.com/spf13/cobra"
)

// newPsCmd creates the 'ps' command to list running containers.
func newPsCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ps",
		Short:   "Show running containers",
		Example: "dive ps",
		RunE: func(cmd *cobra.Command, args []string) error {
			containers, err := docker.ListContainers(ctx)
			if err != nil {
				return output.FriendlyError("Could not list containers.", err)
			}
			output.PrintContainerTable(containers)
			return nil
		},
	}
	return cmd
}

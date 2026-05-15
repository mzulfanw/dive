package cli

import (
	"context"

	"github.com/mzulfanw/dive/internal/docker"
	"github.com/mzulfanw/dive/internal/resolver"
	"github.com/mzulfanw/dive/pkg/output"

	"github.com/spf13/cobra"
)

// newPgCmd creates the 'pg' command for PostgreSQL access.
func newPgCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pg",
		Short:   "Enter psql in PostgreSQL container",
		Example: "dive pg",
		RunE: func(cmd *cobra.Command, args []string) error {
			container, err := resolver.ResolvePostgres(ctx)
			if err != nil {
				return output.FriendlyError("Could not find a running PostgreSQL container.", err)
			}
			return docker.ExecPSQL(ctx, container)
		},
	}
	return cmd
}

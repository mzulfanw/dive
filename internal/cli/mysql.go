package cli

import (
	"context"

	"github.com/mzulfanw/dive/internal/docker"
	"github.com/mzulfanw/dive/internal/resolver"
	"github.com/mzulfanw/dive/pkg/output"

	"github.com/spf13/cobra"
)

// newMysqlCmd creates the 'mysql' command for MySQL CLI access.
func newMysqlCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mysql",
		Short:   "Enter mysql client in MySQL container",
		Example: "dive mysql",
		RunE: func(cmd *cobra.Command, args []string) error {
			container, err := resolver.ResolveMySQL(ctx)
			if err != nil {
				return output.FriendlyError("Could not find a running MySQL container.", err)
			}
			return docker.ExecMySQL(ctx, container)
		},
	}
	return cmd
}

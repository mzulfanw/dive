package cli

import (
	"context"

	"github.com/mzulfanw/dive/internal/docker"
	"github.com/mzulfanw/dive/internal/resolver"
	"github.com/mzulfanw/dive/pkg/output"

	"github.com/spf13/cobra"
)

// newRedisCmd creates the 'redis' command for Redis CLI access.
func newRedisCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "redis",
		Short:   "Enter redis-cli in Redis container",
		Example: "dive redis",
		RunE: func(cmd *cobra.Command, args []string) error {
			container, err := resolver.ResolveRedis(ctx)
			if err != nil {
				return output.FriendlyError("Could not find a running Redis container.", err)
			}
			return docker.ExecRedisCLI(ctx, container)
		},
	}
	return cmd
}

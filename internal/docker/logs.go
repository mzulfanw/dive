package docker

import (
	"context"
	"github.com/mzulfanw/dive/internal/types"
	"os/exec"
)

// StreamLogs streams logs from the given container.
func StreamLogs(ctx context.Context, c types.Container) error {
	cmd := exec.CommandContext(ctx, "docker", "logs", "-f", c.ID)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

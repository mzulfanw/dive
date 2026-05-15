package docker

import (
	"context"
	"os"
	"os/exec"

	"github.com/mzulfanw/dive/internal/types"
)

// ExecPSQL enters psql in the given container.
func ExecPSQL(ctx context.Context, c types.Container) error {
	cmd := exec.CommandContext(ctx, "docker", "exec", "-it", c.ID, "psql", "-U", "postgres")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ExecRedisCLI enters redis-cli in the given container.
func ExecRedisCLI(ctx context.Context, c types.Container) error {
	cmd := exec.CommandContext(ctx, "docker", "exec", "-it", c.ID, "redis-cli")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ExecMySQL enters mysql client in the given container.
func ExecMySQL(ctx context.Context, c types.Container) error {
	cmd := exec.CommandContext(ctx, "docker", "exec", "-it", c.ID, "mysql", "-u", "root", "-p")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ExecShell enters a shell in the given container, preferring bash then sh.
func ExecShell(ctx context.Context, c types.Container) error {
	bashCmd := exec.CommandContext(ctx, "docker", "exec", "-it", c.ID, "bash")
	if err := bashCmd.Run(); err == nil {
		return nil
	}
	shCmd := exec.CommandContext(ctx, "docker", "exec", "-it", c.ID, "sh")
	return shCmd.Run()
}

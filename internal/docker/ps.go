package docker

import (
	"context"
	"encoding/json"
	"os/exec"

	"github.com/mzulfanw/dive/internal/types"
)

// ListContainers returns a slice of running containers.
func ListContainers(ctx context.Context) ([]types.Container, error) {
	cmd := exec.CommandContext(ctx, "docker", "ps", "--format", "{{json .}}")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var containers []types.Container
	for _, line := range splitLines(string(out)) {
		if line == "" {
			continue
		}
		var c types.Container
		if err := json.Unmarshal([]byte(line), &c); err == nil {
			containers = append(containers, c)
		}
	}
	return containers, nil
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i, c := range s {
		if c == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

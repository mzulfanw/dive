package resolver

import (
	"context"
	"github.com/mzulfanw/dive/internal/docker"
	"github.com/mzulfanw/dive/internal/types"
	"github.com/mzulfanw/dive/internal/utils"
	"errors"
)

// ResolveShell finds a container by fuzzy name or image.
func ResolveShell(ctx context.Context, name string) (types.Container, error) {
	containers, err := docker.ListContainers(ctx)
	if err != nil {
		return types.Container{}, err
	}
	for _, c := range containers {
		if utils.FuzzyMatch(c.Names, []string{name}) || utils.FuzzyMatch(c.Image, []string{name}) {
			return c, nil
		}
	}
	return types.Container{}, errors.New("no matching container found")
}

package resolver

import (
	"context"
	"github.com/mzulfanw/dive/internal/docker"
	"github.com/mzulfanw/dive/internal/types"
	"github.com/mzulfanw/dive/internal/utils"
	"errors"
)

var pgPatterns = []string{"postgres", "postgresql", "pg", "db"}

// ResolvePostgres finds a running PostgreSQL container.
func ResolvePostgres(ctx context.Context) (types.Container, error) {
	containers, err := docker.ListContainers(ctx)
	if err != nil {
		return types.Container{}, err
	}
	for _, c := range containers {
		if utils.FuzzyMatch(c.Image, pgPatterns) || utils.FuzzyMatch(c.Names, pgPatterns) {
			return c, nil
		}
	}
	return types.Container{}, errors.New("no running PostgreSQL container found")
}

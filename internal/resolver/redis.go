package resolver

import (
	"context"
	"errors"

	"github.com/mzulfanw/dive/internal/docker"
	"github.com/mzulfanw/dive/internal/types"
	"github.com/mzulfanw/dive/internal/utils"
)

var redisPatterns = []string{"redis", "cache"}

// ResolveRedis finds a running Redis container.
func ResolveRedis(ctx context.Context) (types.Container, error) {
	containers, err := docker.ListContainers(ctx)
	if err != nil {
		return types.Container{}, err
	}
	for _, c := range containers {
		if utils.FuzzyMatch(c.Image, redisPatterns) || utils.FuzzyMatch(c.Names, redisPatterns) {
			return c, nil
		}
	}
	return types.Container{}, errors.New("no running Redis container found")
}

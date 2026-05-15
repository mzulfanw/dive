package resolver

import (
	"context"
	"github.com/mzulfanw/dive/internal/docker"
	"github.com/mzulfanw/dive/internal/types"
	"github.com/mzulfanw/dive/internal/utils"
	"errors"
)

var mysqlPatterns = []string{"mysql", "mariadb"}

// ResolveMySQL finds a running MySQL or MariaDB container.
func ResolveMySQL(ctx context.Context) (types.Container, error) {
	containers, err := docker.ListContainers(ctx)
	if err != nil {
		return types.Container{}, err
	}
	for _, c := range containers {
		if utils.FuzzyMatch(c.Image, mysqlPatterns) || utils.FuzzyMatch(c.Names, mysqlPatterns) {
			return c, nil
		}
	}
	return types.Container{}, errors.New("no running MySQL container found")
}

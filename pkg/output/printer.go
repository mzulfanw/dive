package output

import (
	"fmt"

	"github.com/mzulfanw/dive/internal/types"
	"github.com/mzulfanw/dive/internal/utils"
)

// PrintContainerTable prints a table of containers.
func PrintContainerTable(containers []types.Container) {
	fmt.Printf("%-20s %-20s %-20s %-20s\n", "CONTAINER ID", "IMAGE", "NAMES", "STATUS")
	for _, c := range containers {
		fmt.Printf("%-20s %-20s %-20s %-20s\n", c.ID, c.Image, c.Names, c.Status)
	}
}

// FriendlyError prints a user-friendly error and returns it.
func FriendlyError(msg string, err error) error {
	utils.PrintError(fmt.Sprintf("%s\n%s", msg, err.Error()))
	return fmt.Errorf(msg)
}

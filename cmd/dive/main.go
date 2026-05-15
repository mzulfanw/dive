package main

import (
	"context"
	"os"

	"github.com/mzulfanw/dive/internal/cli"
)

// main is the entry point for the dive CLI.
func main() {
	ctx := context.Background()
	rootCmd := cli.NewRootCmd(ctx)
	if err := rootCmd.Execute(); err != nil {
		 os.Exit(1)
	}
}
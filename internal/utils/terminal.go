package utils

import "fmt"

// PrintError prints a formatted error message.
func PrintError(msg string) {
	fmt.Printf("\033[31m%s\033[0m\n", msg)
}

// PrintInfo prints a formatted info message.
func PrintInfo(msg string) {
	fmt.Printf("\033[36m%s\033[0m\n", msg)
}

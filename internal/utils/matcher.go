package utils

import (
	"strings"
)

// FuzzyMatch returns true if the target matches any of the patterns (case-insensitive, substring).
func FuzzyMatch(target string, patterns []string) bool {
	target = strings.ToLower(target)
	for _, p := range patterns {
		if strings.Contains(target, strings.ToLower(p)) {
			return true
		}
	}
	return false
}

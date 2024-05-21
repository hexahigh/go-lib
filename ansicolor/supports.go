package ansicolor

import (
	"os"
	"strings"
)

// Returns true if TrueColor is supported, false otherwise.
// Might be wrong due to not all terminals setting environment variables correctly.
func SupportsTrueColor() bool {
	colorterm := os.Getenv("COLORTERM")
	term := os.Getenv("TERM")

	// Check COLORTERM for explicit TrueColor indication
	if colorterm == "truecolor" || colorterm == "24bit" {
		return true
	}

	// Check TERM for keywords suggesting TrueColor support
	if strings.Contains(term, "truecolor") || strings.Contains(term, "24bit") {
		return true
	}

	// Default to false if no positive indication of TrueColor support is found
	return false
}

func Supports256Colors() bool {
	term := os.Getenv("TERM")

	if term == "xterm-256color" || term == "screen-256color" {
		return true
	}

	return false
}

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

	if term == "xterm-kitty" {
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

	switch term {
	case "xterm-256color", "screen-256color", "xterm-kitty":
		return true
	}

	return false
}

// Package hyperspaced provides utilities to improve the aesthetics of your
// favourite phrases with automatic space insertion (ASI).
package hyperspaced

import (
	"github.com/72636c/hyperspaced/internal/transform"
)

var (
	// Spaced is an alias for One that is retained for backwards compatibility.
	Spaced = One

	// SpacedN is an alias for N that is retained for backwards compatibility.
	SpacedN = N
)

// N inserts n spaces between each character in a string.
func N(str string, n int) string {
	return transform.Chars(str, transform.NewSpace(n))
}

// One inserts a space between each character in a string.
func One(str string) string {
	return N(str, 1)
}

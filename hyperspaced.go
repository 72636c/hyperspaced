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

// Lower inserts a space between each character in a string, and converts to
// lowercase.
func Lower(str string) string {
	return LowerN(str, 1)
}

// LowerN inserts n spaces between each character in a string, and converts to
// lowercase.
func LowerN(str string, n int) string {
	return transform.Chars(str, transform.Lowercase, transform.NewSpace(n))
}

// N inserts n spaces between each character in a string.
func N(str string, n int) string {
	return transform.Chars(str, transform.NewSpace(n))
}

// One inserts a space between each character in a string.
func One(str string) string {
	return N(str, 1)
}

// Sponge inserts a space between each character in a string, and alternates
// between lowercasing and uppercasing.
func Sponge(str string) string {
	return SpongeN(str, 1)
}

// SpongeN inserts n spaces between each character in a string, and alternates
// between lowercasing and uppercasing.
func SpongeN(str string, n int) string {
	return transform.Chars(str, transform.Spongecase, transform.NewSpace(n))
}

// Upper inserts a space between each character in a string, and converts to
// uppercase.
func Upper(str string) string {
	return UpperN(str, 1)
}

// UpperN inserts n spaces between each character in a string, and converts to
// uppercase.
func UpperN(str string, n int) string {
	return transform.Chars(str, transform.Uppercase, transform.NewSpace(n))
}

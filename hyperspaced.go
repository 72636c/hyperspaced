package hyperspaced

import (
	"github.com/72636c/hyperspaced/internal/text/transform"
)

// Spaced inserts a space between each character in a string.
func Spaced(str string) string {
	return transform.Spaced(str)
}

// SpacedN inserts n spaces between each character in a string.
func SpacedN(str string, n int) string {
	return transform.SpacedN(str, n)
}

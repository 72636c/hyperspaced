package transform

import (
	"strings"
)

var (
	_ Substring = Lowercase
	_ Substring = Spongecase
	_ Substring = Uppercase
)

// Lowercase converts each substring to lowercase.
func Lowercase(_, _ int, str string) string {
	return strings.ToLower(str)
}

// Spongecase alternates between lowercasing and uppercasing each substring.
func Spongecase(_, index int, str string) string {
	if index%2 == 0 {
		return strings.ToLower(str)
	}

	return strings.ToUpper(str)
}

// Uppercase converts each substring to uppercase.
func Uppercase(_, _ int, str string) string {
	return strings.ToUpper(str)
}

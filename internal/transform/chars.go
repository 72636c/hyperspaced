package transform

import (
	"strings"

	"github.com/72636c/hyperspaced/internal/text"
)

// Substring is a function that transforms a substring during iteration over a
// larger piece of text. The substring usually represents one character, but
// this does not hold when working with transformations that output multiple
// characters per input character.
type Substring func(length int, index int, str string) string

// Chars executes substring transformations on each character in a string.
func Chars(str string, fs ...Substring) string {
	chars := text.SplitString(str)

	length := len(chars)

	var builder strings.Builder

	for index, char := range chars {
		str := char.String()

		for _, f := range fs {
			str = f(length, index, str)
		}

		builder.WriteString(str)
	}

	return builder.String()
}

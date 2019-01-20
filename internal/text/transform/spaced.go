package transform

import (
	"errors"
	"strings"

	"github.com/72636c/hyperspaced/internal/text"
)

// Spaced inserts a space between each character in a string.
func Spaced(str string) string {
	return SpacedN(str, 1)
}

// SpacedN inserts n spaces between each character in a string.
func SpacedN(str string, n int) string {
	if n < 0 {
		panic(errors.New("hyperspaced/internal/text.SpacedN: really?"))
	}

	chars := text.SplitString(str)
	sep := toSeparator(n)

	var builder strings.Builder

	for index, char := range chars {
		builder.WriteString(char.String())

		if index < len(chars)-1 {
			builder.WriteString(sep)
		}
	}

	return builder.String()
}

func toSeparator(length int) string {
	var builder strings.Builder

	for index := 0; index < length; index++ {
		builder.WriteString(" ")
	}

	return builder.String()
}

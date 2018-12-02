package transform

import (
	"errors"

	"github.com/72636c/hyperspaced/internal/text"
)

// Spaced inserts a space between each character in a string.
func Spaced(str string) string {
	return SpacedN(str, 1)
}

// SpacedN inserts n spaces between each character in a string.
func SpacedN(str string, n int) string {
	if n < 0 {
		panic(errors.New("hyperspaced.SpacedN: really?"))
	}

	chars := text.SplitString(str)
	sep := toSeparator(n)
	str = "" // lol variable reassignment

	for _, char := range chars {
		str += char.String() + sep
	}

	if len(str) == 0 {
		return ""
	}

	return str[:len(str)-len(sep)]
}

func toSeparator(length int) (sep string) {
	for index := 0; index < length; index++ {
		sep += " "
	}

	return
}

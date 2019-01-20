package transform

import (
	"errors"
	"strings"
)

// NewSpace constructs a transformation that appends n spaces to each non-final
// character.
func NewSpace(n int) TransformSubstring {
	if n < 0 {
		panic(errors.New("hyperspaced/internal/text/transform.NewSpace: really?"))
	}

	sep := toSeparator(n)

	return func(length, index int, str string) string {
		if index < length-1 {
			return str + sep
		}

		return str
	}
}

func toSeparator(length int) string {
	var builder strings.Builder

	for index := 0; index < length; index++ {
		builder.WriteString(" ")
	}

	return builder.String()
}

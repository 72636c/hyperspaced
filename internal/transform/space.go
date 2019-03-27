package transform

import (
	"errors"
	"strings"
)

// NewSpace constructs a transformation that appends n spaces to each non-final
// substring.
func NewSpace(n int) Substring {
	if n < 0 {
		panic(errors.New("hyperspaced/internal/text/transform.NewSpace: really?"))
	}

	sep := strings.Repeat(" ", n)

	return func(length, index int, str string) string {
		if index < length-1 {
			return str + sep
		}

		return str
	}
}

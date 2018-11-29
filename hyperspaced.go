package hyperspaced

import (
	"github.com/72636c/hyperspaced/internal/transform"
)

func Spaced(str string) string {
	return transform.Spaced(str)
}

func SpacedN(str string, n int) string {
	return transform.SpacedN(str, n)
}

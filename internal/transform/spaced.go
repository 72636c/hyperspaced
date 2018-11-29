package transform

import (
	"golang.org/x/text/unicode/norm"
)

func Spaced(str string) string {
	return SpacedN(str, 1)
}

func SpacedN(str string, n int) string {
	if str == "" {
		return ""
	}

	sep := spaces(n)
	out := ""

	var iter norm.Iter
	iter.InitString(norm.NFC, str)

	for !iter.Done() {
		out += string(iter.Next()) + sep
	}

	return out[:len(str)-1]
}

func spaces(count int) string {
	out := ""

	for index := 0; index < count; index++ {
		out += " "
	}

	return out
}

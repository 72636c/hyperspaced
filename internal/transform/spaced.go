package transform

import (
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

func Spaced(str string) string {
	return SpacedN(str, 1)
}

func SpacedN(str string, n int) string {
	var iter norm.Iter
	iter.InitString(norm.NFC, str)

	var current, out string
	separator := toSeparator(n)

	for !iter.Done() {
		previous := current
		current = string(iter.Next())

		if shouldSeparate(previous, current) {
			out += separator
		}

		out += current
	}

	return out
}

func toSeparator(length int) string {
	out := ""

	for index := 0; index < length; index++ {
		out += " "
	}

	return out
}

func shouldSeparate(previous, current string) bool {
	return !isEmpty(previous) &&
		!isJoinControl(previous) &&
		!isJoinControl(current) &&
		!isModifierSymbol(current) &&
		!isVariationSelector(current)
}

func isEmpty(str string) bool {
	return str == ""
}

func isJoinControl(str string) bool {
	return unicode.In(firstRune(str), unicode.Join_Control)
}

func isModifierSymbol(str string) bool {
	return unicode.In(firstRune(str), unicode.Sk)
}

// http://www.unicode.org/charts/PDF/UFE00.pdf
func isVariationSelector(str string) bool {
	return unicode.In(firstRune(str), unicode.Variation_Selector)
}

func firstRune(str string) rune {
	r, _ := utf8.DecodeRuneInString(str)
	return r
}

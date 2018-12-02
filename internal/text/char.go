package text

import (
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// Char represents a character. Whatever that means.
type Char []rune

func (char Char) String() (str string) {
	for _, r := range char {
		str += string(r)
	}

	return
}

// SplitString splits a string into a slice of characters.
//
// It's obviously not perfect; don't @ me.
func SplitString(str string) []Char {
	var iter norm.Iter
	iter.InitString(norm.NFC, str)

	chars := make([]Char, 0)
	runes := make([]rune, 0)

	for !iter.Done() {
		r, _ := utf8.DecodeRune(iter.Next())

		if areSeparate(runes, r) {
			chars = append(chars, Char(runes))
			runes = []rune{r}
		} else {
			runes = append(runes, r)
		}
	}

	if len(runes) == 0 {
		return chars
	}

	return append(chars, Char(runes))
}

func areSeparate(runes []rune, next rune) bool {
	if len(runes) == 0 {
		return false
	}

	previous := runes[len(runes)-1]

	return !isNull(previous) &&
		!isJoinControl(previous) &&
		!isJoinControl(next) &&
		!isModifierSymbol(next) &&
		!isVariationSelector(next)
}

// https://www.fileformat.info/info/unicode/char/0000/index.htm
func isNull(r rune) bool {
	return r == '\u0000'
}

// http://unicode.org/reports/tr44/#Join_Control
func isJoinControl(r rune) bool {
	return unicode.In(r, unicode.Join_Control)
}

// http://www.fileformat.info/info/unicode/category/Sk/list.htm
func isModifierSymbol(r rune) bool {
	return unicode.In(r, unicode.Sk)
}

// http://www.unicode.org/charts/PDF/UFE00.pdf
func isVariationSelector(r rune) bool {
	return unicode.In(r, unicode.Variation_Selector)
}

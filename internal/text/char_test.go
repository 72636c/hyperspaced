package text_test

import (
	"reflect"
	"testing"

	"github.com/72636c/hyperspaced/internal/text"
)

func Test_SplitString(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    []text.Char
	}{
		{
			description: "Empty String",
			input:       "",
			expected:    []text.Char{},
		},
		{
			description: "One Letter",
			input:       "f",
			expected:    []text.Char{{'f'}},
		},
		{
			description: "Two Letters",
			input:       "hi",
			expected:    []text.Char{{'h'}, {'i'}},
		},
		{
			description: "Accent Normalisation",
			input:       "e\u0301",
			expected:    []text.Char{{'é'}},
		},
		{
			description: "Separate Emoji",
			input:       "\U0001f30f\u2696",
			expected:    []text.Char{{'🌏'}, {'⚖'}},
		},
		{
			// https://emojipedia.org/family-man-light-skin-tone-woman-light-skin-tone-girl-light-skin-tone-baby-light-skin-tone/
			description: "Sequenced Emoji",
			input:       "\U0001f468\U0001f3fb\u200d\U0001f469\U0001f3fb\u200d\U0001f467\U0001f3fb\u200d\U0001f476\U0001f3fb",
			expected:    []text.Char{{'👨', '🏻', '\u200d', '👩', '🏻', '\u200d', '👧', '🏻', '\u200d', '👶', '🏻'}},
		},
		{
			// https://emojipedia.org/female-sleuth/
			description: "Female Variant Selector",
			input:       "\U0001f575\ufe0f\u200d\u2640\ufe0f",
			expected:    []text.Char{{'🕵', '\ufe0f', '\u200d', '♀', '\ufe0f'}},
		},
		{
			// https://emojipedia.org/female-technologist/
			description: "Join Control",
			input:       "\U0001f469\u200d\U0001f4bb",
			expected:    []text.Char{{'👩', '\u200d', '💻'}},
		},
		{
			// https://emojipedia.org/man-type-6/
			description: "Skin Tone Modifier",
			input:       "\U0001f468\U0001f3ff",
			expected:    []text.Char{{'👨', '🏿'}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := text.SplitString(testCase.input)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("expected %+v, received %+v", testCase.expected, actual)
			}
		})
	}
}

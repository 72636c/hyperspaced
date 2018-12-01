package transform_test

import (
	"testing"

	"github.com/72636c/hyperspaced"
)

func Test_Spaced(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "Empty String",
			input:       "",
			expected:    "",
		},
		{
			description: "One Word",
			input:       "AESTHETIC",
			expected:    "A E S T H E T I C",
		},
		{
			description: "Two Words",
			input:       "no u",
			expected:    "n o   u",
		},
		{
			description: "Accent Normalisation",
			input:       "e\u0301",
			expected:    "é",
		},
		{
			description: "Standalone Emoji",
			input:       "\U0001f30f planet scale \u2696",
			expected:    "🌏   p l a n e t   s c a l e   ⚖",
		},
		{
			// https://emojipedia.org/family-man-light-skin-tone-woman-light-skin-tone-girl-light-skin-tone-baby-light-skin-tone/
			description: "Sequenced Emoji",
			input:       "\U0001f468\U0001f3fb\u200d\U0001f469\U0001f3fb\u200d\U0001f467\U0001f3fb\u200d\U0001f476\U0001f3fb",
			expected:    "👨🏻‍👩🏻‍👧🏻‍👶🏻",
		},
		{
			// https://emojipedia.org/female-sleuth/
			description: "Female Variant Selector",
			input:       "\U0001f575\ufe0f\u200d\u2640\ufe0f",
			expected:    "🕵️‍♀️",
		},
		{
			// https://emojipedia.org/female-technologist/
			description: "Join Control",
			input:       "\U0001f469\u200d\U0001f4bb",
			expected:    "👩‍💻",
		},
		{
			// https://emojipedia.org/man-type-6/
			description: "Skin Tone Modifier",
			input:       "\U0001f468\U0001f3ff",
			expected:    "👨🏿",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := hyperspaced.Spaced(testCase.input)
			if actual != testCase.expected {
				t.Errorf(
					"expected %+q (%s), received %+q (%s)",
					testCase.expected,
					testCase.expected,
					actual,
					actual,
				)
			}
		})
	}
}

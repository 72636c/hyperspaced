package transform_test

import (
	"testing"

	"github.com/72636c/hyperspaced/internal/transform"
)

func Test_Spaced(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    string
	}{
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
			description: "Emoji",
			input:       "🌏 planet scale ⚖",
			expected:    "🌏   p l a n e t   s c a l e   ⚖",
		},
		{
			description: "Diversity",
			input:       "👨🏻‍💻",
			expected:    "👨🏻‍💻",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := transform.Spaced(testCase.input)
			if actual != testCase.expected {
				t.Errorf("expected '%s', received '%s'", testCase.expected, actual)
			}
		})
	}
}

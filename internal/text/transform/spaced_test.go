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
			description: "Emojified Buzzword",
			input:       "\U0001f30f planet scale \u2696",
			expected:    "🌏   p l a n e t   s c a l e   ⚖",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := hyperspaced.Spaced(testCase.input)
			if actual != testCase.expected {
				t.Errorf(
					"expected %+[1]q (%[1]s), received %+[2]q (%[2]s)",
					testCase.expected,
					actual,
				)
			}
		})
	}
}

func Test_SpacedN(t *testing.T) {
	testCases := []struct {
		description string
		inputString string
		inputN      int
		expected    string
	}{
		{
			description: "One Space",
			inputString: "AESTHETIC",
			inputN:      1,
			expected:    "A E S T H E T I C",
		},
		{
			description: "Two Spaces",
			inputString: "AESTHETIC",
			inputN:      2,
			expected:    "A  E  S  T  H  E  T  I  C",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := hyperspaced.SpacedN(testCase.inputString, testCase.inputN)
			if actual != testCase.expected {
				t.Errorf(
					"expected %+[1]q (%[1]s), received %+[2]q (%[2]s)",
					testCase.expected,
					actual,
				)
			}
		})
	}
}

func Test_SpacedN_Panics_OnNegativeN(t *testing.T) {
	expectedMessage := "hyperspaced.SpacedN: really?"

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected recovery")
			t.FailNow()
		}

		err, ok := r.(error)
		if !ok {
			t.Errorf("expected error, received %T", r)
			t.FailNow()
		}

		message := err.Error()
		if message != expectedMessage {
			t.Errorf("expected %s, got %s", expectedMessage, message)
		}
	}()

	hyperspaced.SpacedN("", -1)

	t.Error("expected panic")
}

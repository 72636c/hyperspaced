package hyperspaced_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/72636c/hyperspaced"
)

func ExampleN() {
	output := hyperspaced.N("AESTHETIC", 2)

	fmt.Println(output)
	// Output: A  E  S  T  H  E  T  I  C
}

func ExampleOne() {
	output := hyperspaced.One("AESTHETIC")

	fmt.Println(output)
	// Output: A E S T H E T I C
}

func Suite_N(t *testing.T, f func(string, int) string) {
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
			actual := f(testCase.inputString, testCase.inputN)
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

func Suite_N_Panics_OnNegativeN(t *testing.T, f func(string, int) string) {
	expectedMessage := "really?"

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
		if !strings.Contains(message, expectedMessage) {
			t.Errorf("expected substring %s, got %s", expectedMessage, message)
		}
	}()

	f("", -1)

	t.Error("expected panic")
}

func Suite_One(t *testing.T, f func(string) string) {
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
			actual := f(testCase.input)
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

func Test_N(t *testing.T) {
	Suite_N(t, hyperspaced.N)
}

func Test_N_Panics_OnNegativeN(t *testing.T) {
	Suite_N_Panics_OnNegativeN(t, hyperspaced.N)
}

func Test_One(t *testing.T) {
	Suite_One(t, hyperspaced.One)
}

func Test_Spaced(t *testing.T) {
	Suite_One(t, hyperspaced.Spaced)
}

func Test_SpacedN(t *testing.T) {
	Suite_N(t, hyperspaced.SpacedN)
}

func Test_SpacedN_Panics_OnNegativeN(t *testing.T) {
	Suite_N_Panics_OnNegativeN(t, hyperspaced.SpacedN)
}

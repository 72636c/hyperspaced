package hyperspaced_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/72636c/hyperspaced"
)

type TestCase struct {
	input    string
	f        func(string) string
	expected string
}

func ExampleLower() {
	output := hyperspaced.Lower("Millions")

	fmt.Println(output)
	// Output: m i l l i o n s
}

func ExampleLowerN() {
	output := hyperspaced.LowerN("Millions", 2)

	fmt.Println(output)
	// Output: m  i  l  l  i  o  n  s
}

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

func ExampleSponge() {
	output := hyperspaced.Sponge("Millions")

	fmt.Println(output)
	// Output: m I l L i O n S
}

func ExampleSpongeN() {
	output := hyperspaced.SpongeN("Millions", 2)

	fmt.Println(output)
	// Output: m  I  l  L  i  O  n  S
}

func ExampleUpper() {
	output := hyperspaced.Upper("Millions")

	fmt.Println(output)
	// Output: M I L L I O N S
}

func ExampleUpperN() {
	output := hyperspaced.UpperN("Millions", 2)

	fmt.Println(output)
	// Output: M  I  L  L  I  O  N  S
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

func Suite_TestCase(t *testing.T, testCase *TestCase) {
	actual := testCase.f(testCase.input)
	if actual != testCase.expected {
		t.Errorf(
			"expected %s, received %s",
			testCase.expected,
			actual,
		)
	}
}

func Test_Lower(t *testing.T) {
	testCase := &TestCase{
		input:    "MILLIONS of people",
		f:        hyperspaced.Lower,
		expected: "m i l l i o n s   o f   p e o p l e",
	}

	Suite_TestCase(t, testCase)
}

func Test_LowerN(t *testing.T) {
	testCase := &TestCase{
		input:    "MILLIONS of people",
		f:        func(str string) string { return hyperspaced.LowerN(str, 2) },
		expected: "m  i  l  l  i  o  n  s     o  f     p  e  o  p  l  e",
	}

	Suite_TestCase(t, testCase)
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

func Test_Sponge(t *testing.T) {
	testCase := &TestCase{
		input:    "MILLIONS of people",
		f:        hyperspaced.Sponge,
		expected: "m I l L i O n S   O f   p E o P l E",
	}

	Suite_TestCase(t, testCase)
}

func Test_SpongeN(t *testing.T) {
	testCase := &TestCase{
		input:    "MILLIONS of people",
		f:        func(str string) string { return hyperspaced.SpongeN(str, 2) },
		expected: "m  I  l  L  i  O  n  S     O  f     p  E  o  P  l  E",
	}

	Suite_TestCase(t, testCase)
}

func Test_Upper(t *testing.T) {
	testCase := &TestCase{
		input:    "MILLIONS of people",
		f:        hyperspaced.Upper,
		expected: "M I L L I O N S   O F   P E O P L E",
	}

	Suite_TestCase(t, testCase)
}

func Test_UpperN(t *testing.T) {
	testCase := &TestCase{
		input:    "MILLIONS of people",
		f:        func(str string) string { return hyperspaced.UpperN(str, 2) },
		expected: "M  I  L  L  I  O  N  S     O  F     P  E  O  P  L  E",
	}

	Suite_TestCase(t, testCase)
}

package cmd_test

import (
	"strings"
	"testing"

	"github.com/72636c/hyperspaced/internal/cmd"
)

func Test_Spaced(t *testing.T) {
	testCases := []struct {
		description string
		args        []string
		input       string
		expected    string
	}{
		{
			description: "unspecified n",
			args:        []string{"spaced"},
			input:       "AESTHETIC",
			expected:    "A E S T H E T I C\n",
		},
		{
			description: "integer n",
			args:        []string{"spaced", "2"},
			input:       "AESTHETIC",
			expected:    "A  E  S  T  H  E  T  I  C\n",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			reader := strings.NewReader(testCase.input)
			writer := new(strings.Builder)

			cmd.Spaced(reader, writer, testCase.args)

			actual := writer.String()
			if actual != testCase.expected {
				t.Errorf("expected %s, received %s", testCase.expected, actual)
			}
		})
	}
}

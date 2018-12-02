package text

import (
	"bufio"
	"fmt"
	"os"
)

// TransformLine is a function that transforms a line of text.
type TransformLine func(string) string

// LineFilter transforms lines of text as they pass from stdin to stdout.
func LineFilter(transform TransformLine) error {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		output := transform(scanner.Text())

		_, err := fmt.Println(output)
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

package transform

import (
	"bufio"
	"fmt"
	"io"
)

// Line is a function that transforms a line of text.
type Line func(string) string

// LineFilter transforms lines of text as they pass from a reader to a writer.
func LineFilter(
	reader io.Reader,
	writer io.Writer,
	transform Line,
) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		output := transform(scanner.Text())

		_, err := fmt.Fprintln(writer, output)
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

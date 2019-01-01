package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/72636c/hyperspaced/internal/text"
	"github.com/72636c/hyperspaced/internal/text/transform"
)

var usage = `spaced [n]
  n int
    number of spaces between each character (default 1)
`

func Spaced(reader io.Reader, writer io.Writer, args []string) {
	n, err := spacesFromArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, usage)
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	transformLine := func(str string) string {
		return transform.SpacedN(str, n)
	}

	err = text.LineFilter(reader, writer, transformLine)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func spacesFromArgs(args []string) (int, error) {
	if len(args) < 2 {
		return 1, nil
	}

	n, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, err
	}

	return n, nil
}

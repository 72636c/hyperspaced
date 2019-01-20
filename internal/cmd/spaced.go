package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/72636c/hyperspaced/internal/transform"
)

// Spaced executes line filters based on the transformations in package
// hyperspaced.
func Spaced(reader io.Reader, writer io.Writer, args []string) {
	config, err := parseConfig(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, spacedUsage())
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	transformLine := func(str string) string {
		return config.transformN(str, config.n)
	}

	err = transform.LineFilter(reader, writer, transformLine)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

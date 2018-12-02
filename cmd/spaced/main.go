package main

import (
	"fmt"
	"os"

	"github.com/72636c/hyperspaced/internal/text"
	"github.com/72636c/hyperspaced/internal/text/transform"
)

func main() {
	err := text.LineFilter(transform.Spaced)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}

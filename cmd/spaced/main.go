package main

import (
	"fmt"
	"os"

	"github.com/72636c/hyperspaced/internal"
	"github.com/72636c/hyperspaced/internal/transform"
)

func main() {
	err := internal.LineFilter(transform.Spaced)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}

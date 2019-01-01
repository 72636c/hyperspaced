package main

import (
	"os"

	"github.com/72636c/hyperspaced/internal/cmd"
)

func main() {
	cmd.Spaced(os.Stdin, os.Stdout, os.Args)
}

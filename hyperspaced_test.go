package hyperspaced_test

import (
	"fmt"

	"github.com/72636c/hyperspaced"
)

func ExampleSpaced() {
	output := hyperspaced.Spaced("AESTHETIC")

	fmt.Println(output)
	// Output: A E S T H E T I C
}

func ExampleSpacedN() {
	output := hyperspaced.SpacedN("AESTHETIC", 2)

	fmt.Println(output)
	// Output: A  E  S  T  H  E  T  I  C
}

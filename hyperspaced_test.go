package hyperspaced_test

import (
	"fmt"

	"github.com/72636c/hyperspaced"
)

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

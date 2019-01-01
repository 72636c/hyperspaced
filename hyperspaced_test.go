package hyperspaced

import (
	"fmt"
)

func ExampleSpaced() {
	output := Spaced("AESTHETIC")

	fmt.Println(output)
	// Output: A E S T H E T I C
}

func ExampleSpacedN() {
	output := SpacedN("AESTHETIC", 2)

	fmt.Println(output)
	// Output: A  E  S  T  H  E  T  I  C
}

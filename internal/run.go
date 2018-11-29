package internal

import (
	"bufio"
	"fmt"
	"os"
)

type TransformLine func(string) string

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

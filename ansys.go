package main

import (
	"fmt"
)

func main() {
	abcdTable := generate_unique_tuples(8)

	for _, tuple := range *abcdTable {
		fmt.Printf("%v - XOR Result: %v\n", tuple, tuple.Xor())
	}
}

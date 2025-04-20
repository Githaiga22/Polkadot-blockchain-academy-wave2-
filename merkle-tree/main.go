package main

import (
	"fmt"
	"merkle/functions"
)

func main() {
	data := []string{"A", "B", "C", "D", "E", "F", "G"}

	mountains := functions.BuildMMR(data)

	fmt.Println("Merkle Mountain Range:")
	for i, m := range mountains {
		fmt.Printf("Mountain %d → Root: %s | Leaves: %v\n", i+1, m.Root[:8], m.Leaves)
	}
}

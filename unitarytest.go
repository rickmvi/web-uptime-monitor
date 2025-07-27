package main

import (
	"fmt"
)
func test() {
	// Example of a simple unit test
	result := add(2, 3)
	expected := 5
	if result != expected {
		fmt.Printf("Test failed: expected %d, got %d\n", expected, result)
	} else {
		fmt.Println("Test passed")
	}
}

func add(a int, b int) int {
	return a + b
}
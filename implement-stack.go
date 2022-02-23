package main

import "fmt"

// LIFO(Last In First Out)
// implement a stack structure with pop, append, and print top functionalities
// we can implement a stack using a slice object

func main() {

	// Create
	var stack []string

	// Push
	stack = append(stack, "World!")
	stack = append(stack, "Fellow ")
	stack = append(stack, "My ")
	stack = append(stack, "Hello ")

	for len(stack) > 0 {
		// Print top
		n := len(stack) - 1
		fmt.Print(stack[n])

		// Pop
		stack = stack[:n]
	}

	//Output: Hello World
}

package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "new string"
}

func main() {
	name := "nathan"
	fmt.Println("Memory address of name is:", &name) // Memory address of name is: 0xc0000240a0

	// Store pointer in variable
	m := &name
	// Dereference pointer
	fmt.Println("Value at memory address of m is:", *m)

	// Update name using a pointer
	updateName(m)
	fmt.Println("Updated name:", name) // Updated name: new string
}

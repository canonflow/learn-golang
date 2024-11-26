package main

import "fmt"

func main() {
	// ===== Strings =====
	fmt.Println("===== STRINGS =====")
	var nameOne string = "Nathan"
	var nameTwo = "Garzya"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Test"
	nameThree = "Three"
	fmt.Println(nameOne, nameTwo, nameThree)

	// Shorthand (without var) [only in a function / can't use it outside a function]
	nameFour := "Shorthand"
	fmt.Println(nameFour)

	// ===== Integer =====
	fmt.Println("\n===== INTEGER =====")
	var ageOne int = 10
	var ageTwo = 20
	ageThree := 30

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits and Memory
	var numOne int8 = 25
	var unsignedNum uint = 0
	fmt.Println(numOne, unsignedNum)

	// ===== FLOATS =====
	fmt.Println("\n===== FLOAT =====")
	var floatOne float32 = 3.14
	var floatTwo float64 = 2.13123213123213123 // Has a slightly higher precision than float32
	floatThree := 103.231312312321312          // Default float64
	fmt.Println(floatOne, floatTwo, floatThree)
}

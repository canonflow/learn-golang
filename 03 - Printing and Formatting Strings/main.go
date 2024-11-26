package main

import "fmt"

func main() {
	age := 21
	name := "Nathan"

	// ===== Print =====
	fmt.Println("===== PRINT =====")
	fmt.Print("Hello, ")
	fmt.Print("World!\n")
	fmt.Print("New Line\n")

	// ===== Println =====
	fmt.Println("\n===== PRINTLN =====")
	fmt.Println("hello canonFlow!")
	fmt.Println("goodbye canonFlow!")
	fmt.Println("my age is", age, "and my name is", name)

	// ===== Printf ===== (formatted strings)
	fmt.Println("\n===== PRINTF =====")
	// %_ -> format specifier
	fmt.Printf("My age is %v and my name is %v\n", age, name)
	fmt.Printf("My age is %q and my name is %q\n", age, name) // My age is '\x15' and my name is "Nathan"
	fmt.Printf("Age is of type %T\n", age)                    // Age is of type int
	fmt.Printf("You scored %f points!\n", 225.55)             // You scored 225.550000 points!
	fmt.Printf("You scored %0.1f points!\n", 225.55)          // You scored 225.6 points!

	// ===== Sprintf ===== (save formatted strings)
	fmt.Println("\n===== SPRINTF =====")
	var str = fmt.Sprintf("My age is %v and my name is %v\n", age, name)
	fmt.Println("The saved string is: ", str) // The saved string is:  My age is 21 and my name is Nathan
}

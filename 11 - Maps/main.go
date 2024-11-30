package main

import (
	"fmt"
)

func main() {
	// String as key type
	fmt.Println("===== STRING AS KEY TYPE =====")
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	// Looping maps
	fmt.Println("\n===== LOOPING =====")
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// Integer as key type
	fmt.Println("\n\n===== INTEGER AS KEY TYPE =====")
	phonebook := map[int]string{
		12345678: "nathan",
		87654321: "garzya",
		12382123: "santoso",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[12345678])

	// Change the value
	phonebook[12345678] = "canonflow"
	fmt.Println(phonebook[12345678])
}

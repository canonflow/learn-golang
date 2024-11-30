package main

import (
	"fmt"
)

func updateName(x string) {
	x = "new string"
}

func updateMenu(y map[string]float64) {
	y["chocolate"] = 12.99
}

func main() {
	// Group A types (non-pointer values) -> string, int, bool, float, array, struct
	name := "nathan"
	updateName(name)
	fmt.Println(name) // still 'nathan'

	// Group B types (pointer wrapper values) -> slice, map, function
	menu := map[string]float64{
		"pie":       7.99,
		"chocolate": 5.99,
	}

	updateMenu(menu)
	fmt.Println(menu) // map[chocolate:12.99 pie:7.99]
}

package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v\n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v\n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2.0)
}

func main() {
	//sayGreeting("Nathan")
	//sayGreeting("Garzya")
	//sayBye("Santoso")

	//names := []string{"nathan", "garzya", "santoso", "canonflow"}
	//
	//cycleNames(names, sayGreeting)
	//cycleNames(names, sayBye)

	area := circleArea(7.0)
	fmt.Printf("Circle area with radius of 7.0 is: %.3f\n", area)
}

package main

import (
	"fmt"
)

var score = 99.5

// Run with 'go run main.go greetings.go'
func main() {
	sayHello("Nathan")
	for _, point := range points {
		fmt.Println(point)
	}

	showScore()
}

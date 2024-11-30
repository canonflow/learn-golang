package main

import (
	"fmt"
	"math"
)

// Shape Interface
type shape interface {
	area() float64
	circumf() float64
}

// Structs
type square struct {
	length float64
}

type circle struct {
	radius float64
}

// Square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// Circle methods
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) circumf() float64 {
	return math.Pi * c.radius * 2
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is: %.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, shape := range shapes {
		printShapeInfo(shape)
		fmt.Println("---")
	}
}

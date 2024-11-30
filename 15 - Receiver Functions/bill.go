package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

// Format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%.2f \n", k+":", v)
		total += v
	}

	// -25 add empty space on right
	fs += fmt.Sprintf("%-25v ... %.2f", "Total:", total)

	return fs
	/*
		Bill breakdown:
		pie:                      ... $5.99
		cake:                     ... $3.99
		Total:                    ... 9.98
	*/
}

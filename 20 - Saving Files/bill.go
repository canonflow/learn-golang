package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%.2f \n", k+":", v)
		total += v
	}

	// Add tip
	fs += fmt.Sprintf("%-25v ... $%.2f \n", "Tip:", b.tip)
	total += b.tip

	// Total
	// -25 add empty space on right
	fs += fmt.Sprintf("%-25v ... $%.2f", "Total:", total)

	return fs
	/*
		Bill breakdown:
		pie:                      ... $5.99
		cake:                     ... $3.99
		Total:                    ... 9.98
	*/
}

// Update tip
func (b *bill) updateTip(tip float64) {
	//(*b).tip = tip // not mandatory for struct
	b.tip = tip
}

// Add an item to the bill
func (b *bill) addItem(name string, price float64) {
	//(*b).items[name] = price // not mandatory for struct
	b.items[name] = price
}

// Save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		// Stop the entire program and print the error
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}

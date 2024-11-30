package main

import (
	"fmt"
)

func main() {
	mybill := newBill("Nathan")

	mybill.addItem("Onion Soup", 4.5)
	mybill.addItem("Veg Pie", 8.95)
	mybill.addItem("Toffee Pudding", 4.95)
	mybill.addItem("Coffee", 3.25)
	mybill.updateTip(10.0)

	fmt.Println(mybill.format())
}

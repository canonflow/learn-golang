package main

import (
	"fmt"
)

func main() {
	mybill := newBill("Nathan")

	fmt.Println(mybill.format())
}

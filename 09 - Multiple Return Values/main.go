package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, name := range names {
		initials = append(initials, name[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn, sn := getInitials("nathan garzya")
	fmt.Println(fn, sn) // N G

	fn2, sn2 := getInitials("santoso")
	fmt.Println(fn2, sn2) // S _
}

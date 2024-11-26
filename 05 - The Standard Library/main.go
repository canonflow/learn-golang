package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// ===== STRINGS =====
	greeting := "Hello there friends!"
	fmt.Println("===== STRINGS =====")
	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "there"))
	fmt.Println(strings.Split(greeting, " "))

	// ===== INTEGER =====
	fmt.Println("\n===== SORT =====")
	ages := []int{45, 20, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	// Sort Descending (if u want descending, write ages[i] < ages[j]
	//sort.Slice(ages, func(i, j int) bool {
	//	return ages[i] > ages[j]
	//})
	fmt.Println("Sorted Integer (Ascending)", ages)

	index := sort.SearchInts(ages, 1010) // If the x doesn't exist, it will return len(slice), which in this case is 7
	fmt.Println("Index of 30:", index)

	names := []string{"nathan", "garzya", "santoso", "canon", "flow"}
	sort.Strings(names)
	fmt.Println("Sorted Strings (Ascending)", names)

	stringIndex := sort.SearchStrings(names, "nathan")
	fmt.Println("Index of nathan:", stringIndex)
}

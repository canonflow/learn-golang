package main

import "fmt"

func main() {
	// ===== ARRAY (fixed length) =====
	var ages [3]int = [3]int{20, 25, 30}
	//var ages = [3]int{20, 25, 30}
	names := [4]string{"Nathan", "Garyza", "Santoso", "canonFlow"}
	names[1] = "Garzya" // Change
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// ===== SLICES (dynamic length, use arrays under the hood) =====
	var scores = []int{100, 50, 60}
	scores[2] = 75
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// ===== SLICE RANGES =====
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, len(rangeOne))
	fmt.Println(rangeTwo, len(rangeTwo))
	fmt.Println(rangeThree, len(rangeThree))

	rangeOne = append(rangeOne, "Koopa")
	fmt.Println(rangeOne, len(rangeOne))
}

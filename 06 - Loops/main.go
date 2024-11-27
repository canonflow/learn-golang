package main

import "fmt"

func main() {
	//x := 0
	//for x < 5 {
	//	fmt.Println("Value of x is:", x)
	//	x++
	//}

	//for i := 0; i < 5; i++ {
	//	fmt.Println("Value of x is:", i)
	//}

	names := []string{"nathan", "garzya", "santoso", "canonflow"}
	//for i := 0; i < len(names); i++ {
	//	fmt.Println(names[i])
	//}

	//for index, name := range names {
	//	fmt.Printf("The value at index %v is %v\n", index, name)
	//}

	for _, name := range names {
		fmt.Printf("The value is %v\n", name)
		name = "new String" // Doesn't update the value inside names
	}

	fmt.Println(names)
}

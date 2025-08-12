package main

import (
	"fmt"
	"sort"
)

func test(ID []int) int {
	missingID := -1
	sort.Ints(ID)

	for i := 0; i < len(ID); i++ {
		if ID[i] != i+1 {
			return ID[i] - 1
		}
	}

	return missingID
}

func main() {
	//create hypothetical array with variable input size
	var x int
	fmt.Print("Please enter the number of IDs: ")
	fmt.Scanln(&x)

	if (x<=0) {
		fmt.Println("Invalid input")
	}

	// Create a slice with user-defined size
	slice := make([]int, x)

	// Fill slice with user input
	fmt.Println("Enter", x, "elements:")
	for i := 0; i < x; i++ {
		fmt.Scan(&slice[i])
	}

	// send the array to method
	result := test(slice)

	//print returned output
	fmt.Println("Missing ID: ", result)

}

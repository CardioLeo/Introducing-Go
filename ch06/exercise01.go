package main

import (
	"fmt"
)

func sum(inputs []int) int {
	var newSum int = 0
	for i := 0; i < len(inputs); i++ {
		newSum = newSum + inputs[i]
		fmt.Println(newSum)
	}
	return newSum
}

func main() {
	// mySlice := make([]int, 3, 3)
	var mySlice = []int{3, 6, 9}
	fmt.Println("The numbers we will add are", mySlice)
	mySum := sum(mySlice)
	fmt.Println("The result is", mySum)
}

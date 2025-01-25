package main

import (
	"fmt"
)

func sum(inputs []int) int {
	var newSum int = 0
	for i := 0; i < len(inputs); i++ {
		fmt.Println(inputs[i])
		newSum += inputs[i]
	}
	return newSum
}

func main() {
	mySlice := make([]int, 3, 3)
	mySum := sum(mySlice)
	fmt.Println(mySum)
}

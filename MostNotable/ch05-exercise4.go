package main

import (
	"fmt"
)

func main() {
	x := []int{
		49, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	temp := x[0]
	for i := 0; i < len(x); i++ {
		if (x[i] < temp) {
			temp = x[i]
		}
	}
	fmt.Println("The Lowest Number is:", temp)
}

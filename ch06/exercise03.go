package main

import (
	"fmt"
)

func findGreatestNum(args ...int) int {
	highest := 0
	for _, v := range args {
		if v > highest {
			highest = v
			fmt.Println("The highest is now:", highest)
		}
	}
	return highest
}

func main() {
	thisHighest := findGreatestNum(7, 2, 34, 5, 123, 5, 56278, 23, 945)
	fmt.Println(thisHighest)
}

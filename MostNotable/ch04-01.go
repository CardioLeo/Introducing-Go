package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 100 {
		// fmt.Println(i)
		if (i % 3 == 0) && (i % 5 == 0) {
                        fmt.Println("FizzBuzz")
                } else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
		i = i + 1
	}
}

// transformed this into my first implementation of FizzBuzz
// based on some of the code from the first five pages of chap.4

// updated it so it fulfills the requirements for the third
// exercise of the fourth chapter

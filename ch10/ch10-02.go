package main

import (
	"fmt"
	"time"
	"math/rand"
)

var n2 int = 0

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i, ":", n2)
		n2++
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

package main

import (
	"fmt"
)

func f(x int) {
	fmt.Println(x)
}

var x int = 5

func f2() {
	fmt.Println(x)
}

func f3() int {
	return f4()
}

func f4() int {
	return 1
}

func f5() (r int) {
	r = 1
	return
}

func f6() (int, int) {
	return 5, 6
}

func main() {
	x := 5
	f(x)
	f2()
	fmt.Println(f3())
	fmt.Println(f5())
	x, y := f6()
	fmt.Println(x, y)
}

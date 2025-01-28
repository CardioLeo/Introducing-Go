package main

import (
	"errors"
	"fmt"
	"container/list"
)

func main() {
	err_1 := errors.New("Error message")
	fmt.Println(err_1)

	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	}

}

package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	file, err := os.Create("test2.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	
	// read the file
	file.WriteString("this is my second test text for my second test text file")
	
	bs, err := ioutil.ReadFile("test2.txt")
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

package main

import (
	"fmt"
)

func main() {
	elements := map[string]map[string]string{
		"Li": map[string]string{
			"name":"Lithium",
			"state":"gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}

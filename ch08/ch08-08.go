package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// next follows the second example, the ways of doing it bottom of page 72

type ByAge []Person
func (this ByAge) Len2() int {
	return len(this)
}

func (this ByAge) Less2(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap2(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

}

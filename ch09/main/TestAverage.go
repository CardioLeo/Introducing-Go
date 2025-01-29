package main

import (
        "fmt"
	"testing"
)

func TestAverage(t *testing.T) {
        v := Average([]float64{1, 2})
        if v != 1.5 {
                fmt.Println("Error: Expected 1.5, got ", v)
        } else {
                fmt.Println("No errors found")
        }
}

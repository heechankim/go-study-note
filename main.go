package main

import (
	"fmt"
	"errors"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Cannot use negative integer")
	}
	return math.Sqrt(f), nil
}

func main() {
	if f, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(f)
	}
}
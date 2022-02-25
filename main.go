package main

import (
	"fmt"
	"strconv"
)

func displayInt(s string) {
	if v, err := strconv.Atoi(s); err != nil {
		fmt.Printf("%s is not a integer.\n", s)
	} else {
		fmt.Printf("The value is %d\n", v)
	}
}

func main() {
	displayInt("two")
	displayInt("2")
}

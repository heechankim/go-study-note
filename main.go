package main

import "fmt"

type rect struct {
	w, h float64
}

func main() {
	r1 := rect{1, 2}
	r2 := new(rect)
	r2.w, r2.h = 3, 4
	r3 := &rect{}
	r4 := &rect{5, 6}

	fmt.Printf("%p \n\n", &r1)

	fmt.Printf("%p \n", &r2)
	fmt.Printf("%p \n\n", r2)

	fmt.Printf("%p \n", &r3)
	fmt.Printf("%p \n\n", r3)

	fmt.Printf("%p \n", &r4)
	fmt.Printf("%p \n\n", r4)
}
package main

import "fmt"

func main() {

	var p *int
	var pp **int
	i := 1
	p = &i
	pp = &p
	fmt.Println(i, *p, **pp)

	i += 1
	fmt.Println(i, *p, **pp)

	*p++
	fmt.Println(i, *p, **pp)

	**pp++
	fmt.Println(i, *p, **pp)
}
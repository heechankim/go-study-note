package main

import "fmt"

type Option struct {
	name string
	value string
}

type Item struct {
	name string
	price float64
	quantity int
	Option
}

func main() {
	shoes := Item{"shoes", 30000, 2,
		Option{"color", "red"}}

	fmt.Println(shoes)

	fmt.Println(shoes.name)

	fmt.Println(shoes.Option.name)
	fmt.Println(shoes.value)
}
package main

import "fmt"

type Item struct {
	name string
	price float64
	quan int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quan)
}

func main() {

	shirt := Item{name: "Shirt", price: 25000, quan: 3}

	fmt.Println(shirt.Cost())

}
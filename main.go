package main

import "fmt"

type Item struct {
	name string
	price float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

type DiscountItem struct {
	Item
	discountRate float64
}

func main() {
	shoes := Item{"shoes", 30000, 2}

	eventShoes := DiscountItem {
		Item{"shoes2", 50000, 3},
		10.00,
	}
	fmt.Println(shoes.Cost())
	fmt.Println(eventShoes.Cost())
}
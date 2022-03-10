package main

import "fmt"

type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println("cost: ", c.Cost())
}

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

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

type Rental struct {
	name string
	feePerDay float64
	periodLength int
	RentalPeriod
}
type RentalPeriod int

const (
	Days RentalPeriod = iota
	Weeks
	Months
)

func (p RentalPeriod) ToDays() int {
	switch p {
	case Weeks:
		return 7
	case Months:
		return 30
	default:
		return 1
	}
}

func (r Rental) Cost() float64 {
	return r.feePerDay * float64(r.ToDays() * r.periodLength)
}

type Items []Coster

func (ts Items) Cost() (c float64) {
	for _, t := range ts {
		c += t.Cost()
	}
	return
}


func main() {
	shoes := Item{"shoes", 25000, 3}
	eventShoes := DiscountItem{
		Item{"eventShoes", 50000, 3},
		10.00,
	}
	video := Rental{"movie", 1000, 3, Days}

	items := Items{shoes, eventShoes, video}
	displayCost(items)
}
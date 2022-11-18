# Polymorphism

### Code
```go
package main

import "fmt"

type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println("cost :", c.Cost())
}

type Item struct {
	name  string
	price float64
	stock int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.stock)
}

type Subscribe struct {
	name        string
	hourlyPrice float64
}

func (t Subscribe) Cost() float64 {
	return t.hourlyPrice * 720
}

func main() {
	item1 := Item{"item1", 20000, 2}
	displayCost(item1)

	onlineService := Subscribe{"online1", 100}
	displayCost(onlineService)
}
```
### Result
```
cost : 40000
cost : 72000
```
### Lessons
- 타입은 인터페이스의 is-A 관계
- 아무런 연관이 없는 타입이 인터페이스의 정의대로 동작이 같도록 정의할 수 있음.
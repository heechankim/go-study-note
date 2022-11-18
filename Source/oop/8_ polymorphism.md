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

func main() {
	item1 := Item{"item1", 20000, 2}

	displayCost(item1)
}

```
### Result


### Lessons
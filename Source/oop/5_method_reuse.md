# Method Reuse

### Code
```go
package main

import "fmt"

type Product struct {
	name string
	Detail
}

type Detail struct {
	price int
	stock int
}

func (d Detail) buyAll() int {
	return d.price * d.stock
}

func main() {
	item := Product{"t-shirt", Detail{10000, 5}}
	
	fmt.Println("If you buy all item then price would be", item.buyAll())
}
```
### Result
```
If you buy all item then price would be 50000
```

### Lessons
- Product 구조체 타입에는 buyAll 메서드가 정의되어 있지 않지만, 임베디드 필드로 가진 Detail이 가지고 있으므로 buyAll 메서드를 호출할 수 있음.

---

### Code
```go
package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	name string
	Detail
}

func (p Product) buyAll() string {
	return "If you buy all " + p.name + " then price would be " + strconv.Itoa(p.Detail.buyAll())
}

type Detail struct {
	price int
	stock int
}

func (d Detail) buyAll() int {
	return d.price * d.stock
}

func main() {
	item := Product{"t-shirt", Detail{10000, 5}}

	fmt.Println(item.buyAll())
}
```
### Result
```
If you buy all t-shirt then price would be 50000
```

### Lessons
- 임베디드 필드의 메서드를 오버라이딩(Overriding)할 수 있다.
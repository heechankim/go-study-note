# Interface

### Code
```go
package main

import "fmt"

type shaper interface {
	area() float64
}

func describe(s shaper) {
	fmt.Println("area :", s.area())
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * 3.14
}

func main() {
	r := rect{3, 4}
	describe(r)

	c := circle{3}
	describe(c)
}
```

### Result
```
area : 12
area : 28.26
```
### Lessons
- 인터페이스는 객체의 동작 방식을 추상적으로 정의함.
- 덕 타이핑 방식으로 객체의 변수나 메서드 집합이 객체를 결정함.
- 보통 인터페이스는 메서드이름 + er을 붙여 네이밍하며, 적어도 한개 많게는 세개 정도의 메서드를 포함하는 짧은 단위로 만든다.
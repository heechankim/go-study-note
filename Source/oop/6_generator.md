# Information Hiding

### Code
```go
# Package Product
package product

type Product struct {
	name string
	price int
}

func New(name string, price int) *Product {
	return &Product{name, price}
}
```
```go
package main

import "product"

func main() {
	item := product.New("t-shirt", 10000)
}
```
### Lessons
- Go에서 식별자의 첫 문자가 대문자인지 소문자인지에 따라 Public, Private를 나눔 (정보은닉)
- 따라서 대문자로 시작하는 필드(exported field)는 외부에서 접근 가능하고, 소문자로 시작하는 필드 (non-exported field)는 패키지 내부에서만 접근할 수 있음.
- 일반적으로 생성함수는 New로 지정하여 외부에서 객체를 생성할 수 있도록 함.
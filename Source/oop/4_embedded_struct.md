# Embedded Struct

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

func main() {
	item := Product{"t-shirt", Detail{10000, 5}}

	fmt.Println(item)
	fmt.Println("Item Name :", item.name)
	fmt.Println("Item Price :", item.Detail.price)
	fmt.Println("Item Stock :", item.stock)
}
```

### Result
```
{t-shirt {10000 5}}
Item Name : t-shirt
Item Price : 10000
Item Stock : 5
```

### Lessons
- OOP에서 프로그램의 규모가 커질수록 상속 관계에서 발생하는 여러 문제들을 피하기 위해 Go에서는 조합을 사용.
- 사용자 정의 타입을 조합하여 구조체를 정의하는 방식으로 객체를 재사용 함.
- 사용자 정의 타입을 구조체의 필드로 지정하는 것을 임베딩이라함.
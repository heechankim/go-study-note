# Object

### Code
```go
package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func (p Person) print() string {
	return "Name is " + p.name + " and Age is " + strconv.Itoa(p.age)
}

func main() {
	hong := Person{name: "Hong-GilDong", age: 25}

	fmt.Println(hong.print())
}
```

### Result
```
Name is Hong-GilDong and Age is 25
```

### Lessons
- Go에서는 상태를 표현하는 타입과 동작을 표현하는 메서드를 분리하여 정의.

  - C언어에서 구조체와 함수를 통해 OOP를 구현하는 것과 비슷한 방식.

- 사용자 정의 타입 - 구조체와 인터페이스를 통해 사용자가 타입을 생성할 수 있음.

  - type 타입명 타입명세
  - 타입명세는 Go의 기본 타입, 구조체, 인터페이스, 함수서명을 사용할 수 있음
- 메서드를 (리시버를 통해) 사용자 정의 타입과 함수를 바인딩하는 방식으로 정의
- func (리시버명 리시버타입) 메서드명(매개변수) (리턴 타입/값) {}
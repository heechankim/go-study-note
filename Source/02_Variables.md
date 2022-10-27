# Variables

#### Code
```go
package main

import "fmt"

func main() {
	var num int
	var text string

	num = 10
	text = "Hello World!"

	fmt.Println(num)
	fmt.Println(text)
}

```
#### Result

```
10
Hello World!
```

#### Lessons

- var name type 이러한 변수 선언의 순서는 'Var'iable 'name' is 'type' 처럼 사람이 읽기 편하도록 설계한 것이라고 한다.

- Go의 자료형
  - Basic Type : Numbers, Strings, Booleans 
  - Aggregate Type : Array, Structs
  - Reference Type : Pointers, Slices, Maps, Functions, Channels
  - Interface Type



<hr>

#### Code
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	whatisthis := 3.141592

	fmt.Println(whatisthis, "is type of", reflect.TypeOf(whatisthis))
}

```

#### Result
```
3.141592 is type of float64
```

#### Lessons
- := 키워드를 통해 선언과 동시에 값으로 초기화를 할 수 있음.
- 동적으로 타입이 지정되는 것을 알 수 있다.
- Convention : 변수의 선언방식, type에 신경쓰기 보단 로직에 집중하도록 하기 위해 일반적으로 수명주기가 짦은 변수에 할당한다.
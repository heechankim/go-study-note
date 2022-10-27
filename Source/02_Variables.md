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
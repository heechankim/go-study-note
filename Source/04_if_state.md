# If Statement

#### Code
```go
package main

import "fmt"

func main() {
	var name string

	name = "chan"

	if name != "chan" {
		fmt.Println("Name is not chan")
	} else {
		fmt.Println("Name is chan")
	}
}

```
#### Result
```
Name is chan
```
#### Lessons
- if 문의 조건식에는 소괄호를 사용하지 않는다.
- 논리 연산자를 통해 조건식을 병합할 때는 소괄호를 사용할 수 있다.
- 조건식에는 반드시 booleans 타입이어야 한다.
- if , if-else, if else-if else 등

<hr>

#### Code
```go
package main

import "fmt"

func main() {

	if name := "kim"; name != "chan" {
		fmt.Println("Name is not chan")
	} else {
		fmt.Println("Name is chan")
	}
}

```

#### Result
```
Name is not chan
```
#### Lessons
- 변수에 특정 값을 할당하거나, if 문에서 사용할 변수를 선언하기 위해 if 문의 초기화 구문을 사용한다.
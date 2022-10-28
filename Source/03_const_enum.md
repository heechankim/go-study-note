# Constant Variables And Enumeration

#### Code
```go
package main

import "fmt"

func main() {
	const maxValue = 64

	fmt.Println("The Max Value is", maxValue)

	maxValue = 48

	fmt.Println("The Max Value is", maxValue)
}
```
#### Result
```
./main.go:10:2: cannot assign to maxValue (untyped int constant 64)
```
#### Lessons
- const 키워드로 상수 선언

<hr>

#### Code
```go
package main

import "fmt"

func main() {
	const (
		first  = 0
		second = 1
		third  = 2
	)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}

```
#### Result
```
1
2
3
```
#### Lessions
- Go에서 열거형의 정의 : 1씩 증가하는 상수의 모음
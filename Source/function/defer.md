# Defer Keyword

### Code
```go
package main

import "fmt"

func main() {
	first()
}
func first() {
	fmt.Println("1. First Start")
	defer second()
	fmt.Println("1. First End")
}
func second() {
	fmt.Println("2. Second Deferred")
}
```
### Result
```
1. First Start
1. First End
2. Second Deferred
```
### Lessons
- defer 구문은 함수 내에서 함수가 종료되기 직전까지 구문을 지연시킨뒤 실행시킨다.

---

### Code
```go
package main

import "fmt"

func someIteration() {
	for i := 0; i < 3; i++ {
		defer fmt.Printf("iteration : %d\n", i)
	}
}

func main() {
	someIteration()
}
```
### Result
```
iteration : 2
iteration : 1
iteration : 0
```
### Lessons
- 결과에서 처럼 defer 구문은 가장 처음 defer 를 만난 구문부터 stack 자료구조에 쌓아 두었다가 함수가 종료되면 stack에서 구문을 꺼내어 실행한다.

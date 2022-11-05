# Closure

### Code
```go
package main

import "fmt"

func main() {
	calc := func(a, b int) int {
		return a + b
	}(2, 5)

	fmt.Println("Closure - calc is", calc)
}
```
### Result
```
Closure - calc is 7
```
### Lessons
- 클로저(Closure)는 익명함수.
- Go 언어에서 함수는 일급 객체이므로 변수에 함수를 할당할 수 있다.


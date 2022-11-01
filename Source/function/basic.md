# Function Basic

### Code
```go
package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func main() {
	result := multiply(2, 5)
	fmt.Println("two times five is", result)
}
```
### Result
```
two times five is 10
```
### Lessons
- go 언어에서 함수의 기본적인 골격은 func 함수명(매개변수) 반환형 { 함수 몸체 } 형태를 가짐

---

### Code
```go
package main

import "fmt"

func etc(str string, ints ...int) {
	fmt.Println("Input String is", str)
	for i := 0; i < len(ints); i++ {
		fmt.Println(ints[i])
	}
}

func main() {
	etc("Chan", 4, 3, 2, 1)
}
```
### Result
```
Input String is Chan
4
3
2
1
```
### Lessons
- 매개변수의 마지막 자리에 '...자료형'의 형태로 개수가 정해지지 않은 가변 인자를 배열의 형태로 받음.
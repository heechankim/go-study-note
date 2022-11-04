# Call By Value and Call By Reference

### Code
```go
package main

import "fmt"

func byValue(num int) {
	num = num + 10
	fmt.Println("In byValue function. num is", num)
}

func main() {
	num := 100
	byValue(num)
	fmt.Println("In main function. num is", num)
}
```
### Result
```
In byValue function. num is 110
In main function. num is 100
```
### Lessons
- 값에 의한 호출(Call by Value)은 전달 받은 매개변수의 값을 복사하여 함수 내부로 전달.

---

### Code
```go
package main

import "fmt"

func byRef(num *int) {
	*num = *num + 10
	fmt.Println("In byRef function. num is", *num)
}

func main() {
	num := 100
	byRef(&num)
	fmt.Println("In main function. num is", num)
}
```
### Result
```
In byRef function. num is 110
In main function. num is 110
```
### Lessons
- 참조에 의한 호출(Call by Reference)은 매개변수 전달시 & 연산자를 통해 변수의 메모리 주소를 전달하고, 매개변수 참조시 * 연산자를 통해 포인터 연산을 수행한다.


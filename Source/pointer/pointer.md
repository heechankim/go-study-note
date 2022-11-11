# Pointer

### Code
```go
package main

import "fmt"

func main() {
	var p *int

	value := 1
	p = &value

	fmt.Println("value =", value)
	fmt.Println("&value =", &value)
	fmt.Println("*p =", *p)
	fmt.Println("p =", p)
}
```
### Result
```
value = 1
&value = 0x14000126008
*p = 1
p = 0x14000126008
```
### Lessons
- C 언어의 포인터 사용법과 동일
- 포인터 변수는 타입앞에 *를 붙여 선언한다.
- 변수의 주소값의 참조는 &를 붙여 참조한다.
- 포인터 변수에는 변수의 주소를 저장할 수 있다.
- \* 연산자로 포인터가 참조하는 변수의 값을 참조할 수 있다.
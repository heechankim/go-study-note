# Function Result and Error Return

### Code
```go
package main

import "fmt"

func someLogic(input int) (result int, err any) {
	result = -1

	if input < 0 {
		err = "The input value is negative integer."
		return
	}

	result = input * input
	err = nil
	return
}

func main() {
	res, _ := someLogic(2)

	fmt.Println(res)
}
```
### Result
```
4
```
### Lessons
- 함수의 처리 결과와 에러 여부를 반환하여 에러 값에 따른 이후 로직을 처리할 수 있도록 함.
- 함수의 반환 값에 _ (blank identifier)를 사용하면 반환 값을 받지만, 변수로써 사용하진 않음

  - "> /dev/null 2>&1"처럼 프로그램의 모든 출력값을 버리는 것과 같은 형태로 사용
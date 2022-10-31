# Switch Statement

### Code
```go
package main

import "fmt"

func main() {

	num := 2

	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	}
}
```
### Result
```
2
```
### Lessons
- switch 값, case 조건의 형식을 갖춘다.
- 첫 번재 case 부터 차례로 검사하고 일치하는 조건이 만나면 해당 블럭을 실행하고 종료한다.
- case의 조건에는 값이 오고 콤마로 구분하여 여러 값을 검사할 수 있다.
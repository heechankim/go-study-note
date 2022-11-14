# 사용자 정의 타입

### Code
```go
package main

import "fmt"

func main() {
	type numbers []int

	var n numbers
	for i := 0; i <= 10; i++ {
		n = append(n, i)
	}

	fmt.Println(n)
}
```
### Result
```
[0 1 2 3 4 5 6 7 8 9 10]
```

### Lessons
- 커스텀 타입은 타입 정의에 사용된 기본타입과 동일하게 동작함.
- 커스텀 타입과 기본 타입의 동작방식은 동일하지만 다른 타입으로 인식하므로 상호간 형변환이 필요함.
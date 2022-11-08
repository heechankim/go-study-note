# For-Range Statement

### Code
```go
package main

import "fmt"

func main() {
	elements := []int{10, 1, 20, 2, 30, 3, 40, 4, 50, 5}

	for index, element := range elements {
		fmt.Printf("Index %d has value %d\n", index, element)
	}
}
```
### Result
```
Index 0 has value 10
Index 1 has value 1
Index 2 has value 20
Index 3 has value 2
Index 4 has value 30
Index 5 has value 3
Index 6 has value 40
Index 7 has value 4
Index 8 has value 50
Index 9 has value 5
```
### Lessons
- 배열과 슬라이스 같은 시퀀스 데이터타입에 for...range 루프를 사용할 수 있다. 
- 전체 요소를 순회하며 인덱스와 값을 반환한다.
- 인덱스가 필요 없을 경우, _ 로 값을 받는다.
# For Statement

### Code
```go
package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println("iteration : ", i)
	}
}
```
### Result
```
iteration :  0
iteration :  1
iteration :  2
```
### Lessons
- go 언어에서 반복문은 for 문만 존재함.
- 여타 프로그래밍 언어와 마찬가지로 초기화 구문, 조건문, 후속작업의 구조를 가짐
- for 문 내에 break, continue 구문을 사용할 수 있음.

<hr>

### Code
```go
package main

import "fmt"

func main() {

LOOP:
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			if i == 1 {
				break LOOP
			}

			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
```
### Result
```
i = 0, j = 0
i = 0, j = 1
i = 0, j = 2
```
### Lessons
- for 문 앞에 콜론으로 끝나는 단어는 레이블.
- Break 문을 레이블과 함께 사용하면 유연한 반복문 제어가 가능함.

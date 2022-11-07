# Array

### Code
```go
package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2}
	arr4 := [...]int{1, 2, 3, 4, 5}
	arr5 := [2][2]int{
		{10, 11},
		{20, 21},
	}

	fmt.Printf("%-10T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-10T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-10T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-10T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%-10T %d %v\n", arr5, len(arr5), arr5)
}
```
### Result
```
[5]int     5 [0 0 0 0 0]
[3]int     3 [1 2 3]
[3]int     3 [1 2 0]
[5]int     5 [1 2 3 4 5]
[2][2]int  2 [[10 11] [20 21]]
```
### Lessons

- 배열의 생성방법은 3가지이다.

  - [길이]타입
  - [길이]타입{초기 값}
  - [...]타입{초기 값}

- 배열의 특징은 다음과 같다

  - 생성당시 크기가 명시적으로 정해짐.
  - 값 타입 (Value Type)
  - 값에 의한 호출 - 값 전체를 복사하여 전달
  - 배열 전체에 대한 비교연산자를 수행할 수 있음

- 내장 함수

  - len(배열) 을 통해 배열의 길이를 구함.
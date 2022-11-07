# Slice

### Code
```go
package main

import "fmt"

func main() {
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 0)
	slice5 := make([]int, 2, 5)

	fmt.Printf("%-10T %d %d %v \n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-10T %d %d %v \n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-10T %d %d %v \n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-10T %d %d %v \n", slice4, len(slice4), cap(slice4), slice4)
	fmt.Printf("%-10T %d %d %v \n", slice5, len(slice5), cap(slice5), slice5)
}
```
### Result
```
[]int      0 0 [] 
[]int      0 0 [] 
[]int      3 3 [1 2 3] 
[]int      0 0 [] 
[]int      2 5 [0 0] 
```
### Lessons
- 슬라이스의 생성방법은 다음과 같다.

  - []타입
  - []타입{초기 값}
  - make([]타입, 길이)
  - make([]타입, 길이, 용량)

- 슬라이스의 특징은 다음과 같다.

  - 길이가 가변적이다.
  - 참조 타입이다.
  - 참조에 의한 호출 - 참조 값 (변수의 주소값)만 전달한다.
  - 비교연산자를 사용한 비교가 불가능하다.

- 내장 함수와 사용

  - cap(slice) : 슬라이스의 용량을 구한다.
  - len(slice) : 슬라이스의 길이를 구한다.
  - append(slice, i) : 슬라이스에 새로운 요소를 붙인다.
  - copy(d, t) : 슬라이스 t의 요소를 슬라이스 d에 복사한다.
  - slice = slice[:cap(slice)] : 슬라이스의 길이를 용량만큼 증가시킨다.
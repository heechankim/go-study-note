# Slicing

### Code
```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println("slice[2:4] is", s[2:4])
	fmt.Println("slice[3:] is", s[3:])
	fmt.Println("slice[:3] is", s[:3])
	fmt.Println("slice[:] is", s[:])
}
```
### Result
```
slice[2:4] is [3 4]
slice[3:] is [4 5]
slice[:3] is [1 2 3]
slice[:] is [1 2 3 4 5]
```
### Lessons
- 파이썬의 슬라이싱 연산과 매우 유사

---

# Slice Append

### Code
```go
package main

import "fmt"

func main() {
	s1 := []int{1}
	s2 := []int{2, 3}
	s3 := []int{6, 7, 8, 9, 10}

	s1 = append(s1, s2...)
	fmt.Println("append(s1, s2)", s1)

	s1 = append(s1, 4, 5)
	fmt.Println("append(s1, 4, 5)", s1)

	s1 = append(s1, s3[:2]...)
	fmt.Println("append(s1, s3[:2)", s1)

}
```
### Result
```
append(s1, s2) [1 2 3]
append(s1, 4, 5) [1 2 3 4 5]
append(s1, s3[:2) [1 2 3 4 5 6 7]
```
### Lessons
- append 내장함수에는 첫번째와 두번째 인수로 들어온 슬라이스를 합한 슬라이스를 리턴한다.

---

### Code
```go
package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	for i := 0; i < 9; i++ {
		s = append(s, i)

		fmt.Printf("len: %d, cap: %d, %v\n", len(s), cap(s), s)
	}
}
```
### Result
```
len: 1, cap: 3, [0]
len: 2, cap: 3, [0 1]
len: 3, cap: 3, [0 1 2]
len: 4, cap: 6, [0 1 2 3]
len: 5, cap: 6, [0 1 2 3 4]
len: 6, cap: 6, [0 1 2 3 4 5]
len: 7, cap: 12, [0 1 2 3 4 5 6]
len: 8, cap: 12, [0 1 2 3 4 5 6 7]
len: 9, cap: 12, [0 1 2 3 4 5 6 7 8]
```
### Lessons
- 배열에서 길이와 용량은 항상 같은 크기이지만, 슬라이스에서 길이와 용량은 같지 않을 수 있다.
- 새로운 요소를 append 하기에 용량이 작다면 자동으로 용량을 증가한 슬라이스를 반환한다.

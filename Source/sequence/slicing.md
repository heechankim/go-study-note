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

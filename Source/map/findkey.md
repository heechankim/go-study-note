# Key 로 맵 접근

### Code
```go
package main

import "fmt"

func main() {
	numbers := make(map[string]int)

	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3

	if v, ok := numbers["one"]; ok {
		fmt.Println("one 키가 맵에 존재 :", v)
	} else {
		fmt.Println("one 키가 맵에 존재하지 않음.")
	}
}
```

### Result
```
one 키가 맵에 존재 : 1
```
### Lessons

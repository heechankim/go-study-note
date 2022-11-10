# Map

### Code
```go
package main

import "fmt"

func main() {
	map1 := map[string]int{}
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3

	map2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	map3 := make(map[string]int, 3)
	map3["one"] = 1
	map3["two"] = 2
	map3["three"] = 3

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
}
```
### Result
```
map[one:1 three:3 two:2]
map[one:1 three:3 two:2]
map[one:1 three:3 two:2]
```
### Lessons
- 맵의 생성방식은 다음과 같다.

  - map[키타입]값타입{}
  - map[키타입]값타입{ 초기 값 }
  - make(map[키타입]값타입, [용량])
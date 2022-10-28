# Constant Variables And Enumeration

#### Code
```go
package main

import "fmt"

func main() {
	const maxValue = 64

	fmt.Println("The Max Value is", maxValue)

	maxValue = 48

	fmt.Println("The Max Value is", maxValue)
}
```
#### Result
```
./main.go:10:2: cannot assign to maxValue (untyped int constant 64)
```
#### Lessons
- const 키워드로 상수 선언

<hr>

#### Code
#### Result
#### Lessions
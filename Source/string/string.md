# Rune

### Code
```go
package main

import (
	"fmt"
	"reflect"
	"unicode"
)

func main() {
	var char1 = 'A'

	fmt.Println(unicode.IsLetter(char1))
	fmt.Println(reflect.TypeOf(char1))
	fmt.Println(char1)
	fmt.Printf("%c \n", char1)
}
```
### Result
```
true
int32
65
A 
```
### Lessons
- go 에서는 문자를 표현하는 자료형은 없다.
- 문자의 표현은 utf8 즉 int32로 표현되며, 재정의된 타입 rune으로 표현한다.

---

# String

### Code
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var string1 = "ABC"

	fmt.Println(reflect.TypeOf(string1))
	fmt.Println(reflect.TypeOf(string1[0]))

	for i := 0; i < len(string1); i++ {
		fmt.Printf("%c \n", string1[i])
	}
}
```
### Result
```
string
uint8
A 
B 
C 
```
### Lessons
- 인덱스 연산자를 사용하여 문자열 내부의 바이트에 접근할 수 있다.

---
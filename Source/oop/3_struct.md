# 구조체

### Code
```go
package main

import "fmt"

type Data struct {
	name  string
	value int
}

func main() {
	d := Data{name: "data1", value: 1}
	fmt.Println(d)
}
```
### Result
```
{data1 1}
```
### Lessons
- 정의방식 : type 구조체명 struct { 필드명 필드타입 }
- 필드 타입이 다른 필드명에 대해서는 줄바꿈으로 구분한다.
- 필드 타입이 같은 필드명에 대해서는 콤마로 구분한다.

---

### Code
```go
package main

import "fmt"

type Data struct {
	name  string
	value int
}

func main() {
	d1 := Data{name: "data1", value: 1}
	d2 := &Data{"data2", 2}
	d3 := new(Data)
	d3.name = "data3"
	d3.value = 3

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(*d2)
	fmt.Println(d3)
	fmt.Println(*d3)
}
```
### Result
```
{data1 1}
&{data2 2}
{data2 2}
&{data3 3}
{data3 3}
```
### Lessons
- 구조체 리터럴 생성 : 구조체타입{초기값}
- 구조체 리터럴의 포인터 생성 : &구조체타입{초기값}
- 구조체 포인터 생성 : new(구조체타입)
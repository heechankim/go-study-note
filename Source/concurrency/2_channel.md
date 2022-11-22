# Channel

### Code
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")

	done := make(chan bool)

	go long(done)
	go short(done)

	<-done
	<-done
	fmt.Println("main end")
}

func long(done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("long function done")
	done <- true
}

func short(done chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println("short function done")
	done <- true
}
```
### Result
```
main start
short function done
long function done
main end
```
### Lessons
- 채널의 생성방법
```go
// 변수 선언 후 채널 할당
var ch chan string
ch = make(chan string)

// 선언과 동시에 채널 할당
ch := make(chan string)
```
- 채널을 통한 값 전송

  - 채널에 값이 존재하면 비워질 때까지 전송 할 수 없다.
  - 채널에 값이 존재하지 않는다면 수신할 수 없다.
```go
// 채널에 값 전송
ch <- "message"

// 채널로 부터 값 수신
msg := <- ch
```

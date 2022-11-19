# Goroutine

### Code
```go
package main

import (
	"fmt"
	"time"
)

func shortTime() {
	fmt.Println("ShortTime Start!\t", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("ShortTime End!\t\t", time.Now())
}

func longTime() {
	fmt.Println("LongTime Start!\t\t", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("LongTime End!\t\t", time.Now())
}

func main() {
	fmt.Println("Main Start!\t\t", time.Now())

	go longTime()
	go shortTime()

	time.Sleep(5 * time.Second)
	fmt.Println("Main End!\t\t", time.Now())
}
```
### Result
```
Main Start!              2022-11-19 13:02:28.072391 +0900 KST m=+0.000087709
ShortTime Start!         2022-11-19 13:02:28.072647 +0900 KST m=+0.000343042
LongTime Start!          2022-11-19 13:02:28.072644 +0900 KST m=+0.000340209
ShortTime End!           2022-11-19 13:02:29.073721 +0900 KST m=+1.001413667
LongTime End!            2022-11-19 13:02:31.073566 +0900 KST m=+3.001251667
Main End!                2022-11-19 13:02:33.073738 +0900 KST m=+5.001415709
```
### Lessons
- go 키워드와 함께 함수를 호출하면 함수는 고루틴에서 실행된다. (경량 스레드)
- 실행중인 고루틴이 있어도, 메인함수가 종료되면 프로그램이 종료된다.

  - 메인함수 내에서 고루틴들의 종료를 확인할 수 있는 로직이 필요.

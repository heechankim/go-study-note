package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start main function", time.Now())

	go long()
	go short()

	time.Sleep(5 * time.Second)
	fmt.Println("End main function", time.Now())
}

func long() {
	fmt.Println("-- Start long function", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("-- End long function", time.Now())
}

func short() {
	fmt.Println("-- Start short function", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("-- End short function", time.Now())
}
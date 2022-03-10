package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start main function", time.Now())

	done := make(chan bool)
	go long(done)
	go short(done)

	<-done
	<-done

	fmt.Println("End main function", time.Now())
}

func long(d chan bool) {
	fmt.Println("-- Start long function", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("-- End long function", time.Now())
	d <- true
}

func short(done chan bool) {
	fmt.Println("-- Start short function", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("-- End short function", time.Now())
	done <- true
}
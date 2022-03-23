package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad Call\r\n")
}
func main() {
	fmt.Println("Calling test")
	test()
	fmt.Println("Test Completed")
}
package main

import "fmt"

func f1() {
	fmt.Println("f1 - start")
	defer f2()
	fmt.Println("f1 - end")
}

func f2() {
	fmt.Println("f2 - deferred")
}

func main() {
	f1()
}

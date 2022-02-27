package main

import "fmt"

func f1() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
func main() {
	f1()
}

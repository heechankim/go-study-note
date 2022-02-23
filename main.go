package main

import "fmt"

func main() {
	var a int = 1

	switch a {
	case 1:
		fmt.Println("a is 1")
		fallthrough
	case 2:
		fmt.Println("a is 2")
		fallthrough
	case 3:
		fmt.Println("a is 3")
	}

}

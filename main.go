package main

import "fmt"

func main() {
	numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	fmt.Println("sum: ", sum)
}
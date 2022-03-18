package main

import (
	"fmt"
	"errors"
)

func main() {
	errNotFound := errors.New("Not found error")

	fmt.Println("error: ", errNotFound)
	fmt.Println("error: ", errNotFound.Error())
}
package lib

import "fmt"

var message string

func init() {
	message = "this is lib package"
}

func IsDigit(c int32) bool {
	fmt.Println(message)
	return '0' <= c && c <= '9'
}
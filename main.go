package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

type SqrtError struct {
	time time.Time
	value float64
	message string
}

func (e SqrtError) Error() string {
	return fmt.Sprintf("[%v]ERROR - %s(value: %g)", e.time, e.message, e.value)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, SqrtError{time.Now(), f, "cannot use n value"}
	}
	return math.Sqrt(f), nil
}

func main() {
	if f, err := Sqrt(-1); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(f)
	}
}
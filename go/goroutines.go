package main

import (
	"fmt"
	"time"
)

func main() {
	go f("sub routines")
	go func(message string) {
		f(message)
	}("going")
	f("main routines")

	time.Sleep(time.Second * 10)
}
func f(message string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(message, ":", i)
	}
}

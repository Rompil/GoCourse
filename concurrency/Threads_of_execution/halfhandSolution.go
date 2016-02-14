package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Just main goroutine is running")
	doGoroutine()
	time.Sleep(time.Duration(200 * time.Millisecond)) //this is not an optimal solution of the problem
	fmt.Println("main goroutine is finished")         // we don't know how long the function takes to execute
}
func doGoroutine() {
	go func() {
		fmt.Println("another goroutine is running!")
	}()
}

//

package main

import (
	"fmt"
	"time"
)

func Race() {
	wait := make(chan struct{})
	var n int
	go func() {
		time.Sleep(time.Duration(1 * time.Second))
		n++
		close(wait)
	}()
	n++
	<-wait
	fmt.Println(n) //due to data race the output can be 1 or 2
}

func main() {
	Race()
}

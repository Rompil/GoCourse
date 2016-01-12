package main

//time go run withoutGoroutine.go
import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 12; i++ {
		multi(i)
	}
}

func multi(x int) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d * %d = %d \n", x, i, x*i)
	}
	time.Sleep(time.Duration(100 * time.Millisecond))
}

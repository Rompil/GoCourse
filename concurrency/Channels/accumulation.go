package main

import "fmt"

func main() {
	chan1 := incrementor("s string")
	chan2 := accumulator(chan1)
	for i := range chan2 {
		fmt.Println(i)
	}
}

func incrementor(s string) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- 1
			fmt.Printf("Current iteration is %d \n", i)
		}
		close(ch)
	}()
	return ch
}
func accumulator(ch chan int) chan int {
	c := make(chan int)
	go func() {
		sum := 0
		for i := range ch {
			sum += i
			c <- sum
		}
		close(c)
	}()
	return c
}

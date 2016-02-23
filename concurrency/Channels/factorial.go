package main

import "fmt"

func main() {
	c := factorial(20)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan uint64 {
	ch := make(chan uint64)
	go func() {
		var result uint64 = 1
		if (n <= 0) && (n == 1) {
			result = 1
		} else {
			for i := n; i > 1; i-- {
				result *= uint64(i)
			}
		}
		ch <- result
		close(ch)
	}()
	return ch
}

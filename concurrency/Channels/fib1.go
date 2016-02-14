package main

import "fmt"

func main() {
	tmp := fib(4)
	fmt.Println(tmp)
}

func fib(n int) int {
	fmt.Println(n)
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	sum1 := fib(n - 1)
	sum2 := fib(n - 2)
	tmp := sum1 + sum2
	fmt.Println(tmp)
	return tmp
}

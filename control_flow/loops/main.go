package main

import (
	"fmt"
)

const N int = 5

func main() {
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			//	fmt.Println(i, " ", j)
		}
	}
	///fmt.Println(i) //we don't see i in this scope. I exist only in the loop space
	// So we can use this variable again. Should we do use two vars with the same name in different places?
	i := 0
	for i < N {
		fmt.Println(fibfunc(i))
		i++
	}
}

////need to fix, it doesn't work
func fibfunc(n int) int {
	first := 1
	second := 1

	if n == 0 || n == 1 {
		return 1
	} else {
		result := 0
		i := 2
		for i <= n {
			result = first + second
			second, first = result, second
		}
		return result
	}

}

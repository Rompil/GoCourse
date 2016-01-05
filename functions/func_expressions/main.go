package main

import (
	"fmt"
)

func main() {
	//functional expression
	digfunc := func(indig int) int {
		return indig % 3
	}
	fmt.Printf("%T \n", digfunc)
	df := makedigfunc(2)
	fmt.Printf("%T \n", df)
}

//more complicated case,the function returns a function as a result
func makedigfunc(div int) func(int) int {
	return func(indig int) int {
		return indig % div
	}
}

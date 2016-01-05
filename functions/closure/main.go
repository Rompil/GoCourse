package main

import (
	"fmt"
)

func main() {
	func1 := makeCounter()
	func2 := makeCounter()
	fmt.Println(func1()) //1
	fmt.Println(func1()) //2
	fmt.Println(func2()) //1 again
}

func makeCounter() func() int {
	var x int = 0       // this variable lives in the closure of whole makeCounter function and a generated function
	return func() int { // call it many time. This var store data between calls but only for one generated function
		x++
		return x
	}
}

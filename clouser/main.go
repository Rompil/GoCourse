package main

import "fmt"

func main() {
	incFunc := wrapper()
	for i := 0; i < 100; i++ {
		incFunc()
	}

}

func wrapper() func() int {
	x := 0 //acts like a static variable. I should remember it
	return func() {
		x++
		fmt.Print(x, "\n")
		return x
	}
}

package main

import (
	"fmt"
)

//N- size of 1D array
const N int = 100

func main() {
	//this is an array
	var m [N]int
	//it's an array too
	b := [N]int{1, 3, 2, 5, 6}
	fmt.Println(b)
	m[37] = 42
	for i := range m {
		m[i] = i
	}
	fmt.Println(m)
	var d2 [N][N]int
	for i := range m {
		for j := range m {
			d2[i][j] = i*10 + j
			fmt.Print(d2[i][j], " ")
		}
		fmt.Println("")
	}
}

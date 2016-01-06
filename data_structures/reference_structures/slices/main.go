//slices are reference type and if I want to pass it to the function I should not use with
package main

import (
	"fmt"
)

func main() {
	//do_not_do_this := []int
	//fmt.Print(do_not_do_this)
	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 30, 100)
	fmt.Println(cap(sl))
	fmt.Println(len(sl))
	fmt.Println(sl2)

	s := make([]string, 3, 10)
	s[0] = "zero"
	s[1] = "one"
	s[2] = "two"
	s = append(s, "three", "four")
	c := make([]string, len(s))
	copy(c, s) //deep copying
	c[2] = "три"
	fmt.Print(s)
	fmt.Print(c)
	c = append(c[:2], c[3:]...) //remove element c[2]
	fmt.Print(c)
}

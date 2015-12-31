package main

import (
	"fmt"
)

const (
	one = iota
	two
)

func main() {
	const three = iota
	fmt.Print(one, "\n")
	fmt.Print(two, "\n")
	fmt.Print(three, "\n")
}

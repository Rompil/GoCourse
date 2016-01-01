package main

import "fmt"

func main() {
	//fmt.Print(x) //it won't work
	x := "string"
	fmt.Print(x, "\n")
	fmt.Print(y, "\n")
}

//it works! declaration y after usage in main. All in one main scope
var y int = 42

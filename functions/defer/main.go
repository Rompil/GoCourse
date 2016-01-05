package main

import "fmt"

func main(){
	defer world()
	hello()
}

func hello(){
	fmt.Print("Hello ")
}

func world(){
	fmt.Println("world")
}
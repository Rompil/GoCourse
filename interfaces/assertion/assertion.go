package main

import "fmt"

func main() {
	var str interface{} = "Роман Пилюгин"
	if s, ok := str.(string); ok { ///here we use assertion about the type of str. str must be interface{} type
		fmt.Println(s)
	} else {
		fmt.Println("Value is not a string.")
	}
	var x interface{} = 7
	fmt.Printf("%T \n", x)
	//but we cannot do:
	//fmt.Println((x+3))//invalid operation: x + 3 (mismatched types interface {} and int)
	fmt.Println((x.(int) + 3))
}

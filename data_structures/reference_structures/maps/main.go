//Samples of using maps in Go
package main

import (
	"fmt"
)

//Maps are reference type so then you address to it you already do it by its ref
func main() {
	m := make(map[string]string) //we store key-value pairs both types are
	m["one"] = "один"
	m["three"] = "три"
	m["seven"] = "eleven"
	fmt.Println(m)

	n := map[string]int{"Лена": 600, "Маша": 650, "Оля": 500} //create a map with instant initialisation. Pay attention on the comma in the end of initialization
	fmt.Println(n["Маша"])

	for key, value := range n { //we use range for a map
		fmt.Printf("%v прыгает на %v см в длину \n", key, value)
	}

	v := m["three"]
	fmt.Println(v)
	val, exist := m["six"]  //we can read data from a dic and check for existence
	fmt.Println(val, exist) //val has "zero" value."" is zero value for a string

	delete(m, "seven") //delete value with key as "seven"
	_, existSeven := m["seven"]
	fmt.Println(existSeven)

}

package main

import (
	"fmt"
)

func main() {
	//	fname := "Максим"
	//	fname := "Антон"
	//	fname := "Ольга"
	fname := "Тимур"

	switch {
	case fname == "Антон":
		fmt.Printf("%v -Друг \n", fname)
	case fname == "Максим":
		fmt.Printf("%v -не просто друг, но и ", fname)
		fallthrough
	case fname == "Ольга":
		fmt.Printf("коллега.")
	default:
		fmt.Printf("Я вообще таких не знаю. \n")
	}

}

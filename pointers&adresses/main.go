package main

import (
	"fmt"
)

func zero(z int) {
	fmt.Println("Zero func")
	fmt.Printf("%v is stored at %#v \n", z, &z)
}

func truezero(z *int) {
	fmt.Println("Truezero func")
	fmt.Printf("%v is stored at %#v \n", *z, z)
}

func main() {
	a := 42 //a value

	fmt.Printf("a adress is %#v \n", &a)
	fmt.Print("Set a new value: ")
	fmt.Scan(&a) //get reference
	fmt.Printf("a values is %#v \n", a)

	var b *int = &a //pointer ot a value
	*b = (*b) << 1  //dereference
	fmt.Printf("A NEW value of a is %#v \n", a)
	fmt.Printf("%v is stored at %#v \n ", a, &a)
	zero(a)
	truezero(&a)
	fmt.Printf("%v is stored at %#v \n ", a, &a)
}

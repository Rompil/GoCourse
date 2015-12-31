package main

import (
	"fmt"
)

func zero(z int){
	fmt.Println("Zero func")
	fmt.Println("%v is stored at %#v", z, &z)
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
	fmt.Println("%v is stored at %#v", a, &a)
	zero(a)

}

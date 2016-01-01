package main

import(
	"fmt"
)

func main() {
	var str string
	str, _ = greet("Роман","Пилюгин")
	fmt.Printf(str)

}
//greet has multiple input and output parameters and uses named return- bad practice
//pay attention on braces  in return values
func greet(first, second string) (s string, n int) {
	s = fmt.Sprint(second," ", first) //Sprint- merge strings in one
	n = len(s)
	return
}
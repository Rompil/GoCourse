package main

import "fmt"

func main(){
	var cnl <-chan string = expFunc()
	for i:= range cnl {
		fmt.Println(i)
	}
}

func expFunc() <-chan string{
	ch:= make(chan string)
	go func() {
		for i:= 0 ; i<100; i++{
			ch <- string(i)
		}
		close(ch)
	} ()
	return ch
}

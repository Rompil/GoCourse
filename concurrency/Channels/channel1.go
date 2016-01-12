package main

import "fmt"

func main() {
	var cnl <-chan string = expFunc()
	for i := range cnl { //channels can be used in for loop with range
		fmt.Println(i)
	}
}

func expFunc() <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- string(i)
		}
		close(ch) //mind the close() in a goroutine
	}()
	return ch
}

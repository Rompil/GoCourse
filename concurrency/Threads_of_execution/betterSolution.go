package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main(){
	wg.Add(1)
	fmt.Println("Just main goroutine is running")
	doGoroutine()
	fmt.Println("main goroutine is finished")
	wg.Wait()
}
func doGoroutine(){
	go func(){
		fmt.Println("another goroutine is running!")//still works not so well
		wg.Done()
	}()
}
//
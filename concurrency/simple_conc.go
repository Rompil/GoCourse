package main

import (
	"fmt"
	"sync"
	"time"
	//	"runtime"
)

var wg sync.WaitGroup

func init() { //this is a special function which prepares environment before the main func execution
	//	runtime.GOMAXPROCS(runtime.NumCPU())//explicitly set the number of utilized cores for gorutines
}

func Foo() {
	for i := 0; i < 1000; i++ {
		fmt.Printf("This is Foo__ func on %v iteration\n", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
}
func Bar() {
	for i := 0; i < 1000; i++ {
		fmt.Printf("This is __Bar func on %v iteration\n", i)
		time.Sleep(time.Duration(20 * time.Millisecond))

	}
}

func main() {
	wg.Add(2)
	go Foo()
	go Bar()
	wg.Wait()
}

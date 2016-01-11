package main
//time go run withGoroutine.go

import(
	"fmt"
	"time"
	"sync"
)
var wg = sync.WaitGroup{}
/*
Problem with this code is that we don't control sequence of execution of each goroutines.
We must use channels. See you in the next series...
*/
func main(){
	for i:=1; i<=12; i++{
		wg.Add(1)
		go multi(i)
	}
	wg.Wait()
}

func multi(x int) {
	for i:=1; i<=12; i++{
		fmt.Printf("%d * %d = %d \n",x,i,x*i)
	}
	time.Sleep(time.Duration(100*time.Millisecond))
	wg.Done()
}
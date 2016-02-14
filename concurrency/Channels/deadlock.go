package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1 //this causes dead-lock, code won't execute unless the value in the channel is read
	// it should be done in a go func
	// go func(){
	//   c<-1
	// close(c)
	// }
	fmt.Println(<-c)
}

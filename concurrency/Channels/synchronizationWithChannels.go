package main

import (
	"fmt"
	"time"
)

func PublishNews(n string, delay time.Duration) (wait <-chan struct{}) { //empty struct means the channel is used only as a signal
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("Breaking news: ", n)
		close(ch)
	}()
	return ch
}

func main() {
	wait := PublishNews("Goroutine with the channel", 1*time.Second)
	fmt.Print("Waiting for the news...\n")
	<-wait //channel carries signal
	fmt.Print("Everithing is over...")
}

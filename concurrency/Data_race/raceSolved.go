package main

import (
	"fmt"
	"time"
)

func Race(){
	wait:= make(chan int)
	go func(){
		n:=0
		time.Sleep(time.Duration(1*time.Second))
		n++
		wait <- n
	}()
	n:=<-wait
	n++
	fmt.Println(n)//due to data race the output must be 2
}

func main(){
	Race()
}
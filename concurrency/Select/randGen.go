package main

import (
	"fmt"

)

func main(){
	a:=rand()
	fmt.Println(a[0],a[1])
}

func rand()[]int{
	ch :=make(chan int, 10 )
	go func(){
		for i:=0; i<1; i++{
			select{
				case ch<-0:
				case ch<-1:
			}
		}
		close(ch)
	}()
	return(ch)

var arr []int = rand()
	for i:= range ch{
		arr[i]++
	}
	return arr
}
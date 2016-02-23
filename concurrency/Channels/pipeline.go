package main

// This example implements an idea of pipelining in this article (http://blog.golang.org/pipelines)
// The sequance of stages are applied to numbers.
import "fmt"

func main() {
	c := Consumer(midStage(Producer(1, 2, 3, 4, 5)))
	for i := range c {
		fmt.Println(i)
	}
}

// there maust be a producer function
func Producer(in ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range in {
			ch <- n
		}
		close(ch)
	}()
	return ch
}

//meduim stage
func midStage(in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := range in {
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

// there must be a consumer function
func Consumer(in <-chan int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := range in {
			ch <- fmt.Sprintf("current value is : %d \n", i)
		}
		close(ch)
	}()
	return ch
}

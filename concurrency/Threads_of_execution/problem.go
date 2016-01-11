package main
//Example of a problem with goroutines
import "fmt"

func main(){
	fmt.Println("Just main goroutine is running")
	doGoroutine()
	fmt.Println("main goroutine is finished")
}
//this function doesn't output because the main goroutine finished earlier
func doGoroutine(){
	go func(){
		fmt.Println("another goroutine is running!")
	}()
}
//
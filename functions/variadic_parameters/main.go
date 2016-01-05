package main

import (
	"fmt"
)

func main() {
	fmt.Print("Среднее значение: ")
	data := []float64{43, 6, 132, 65, 6, 7, 987}
	fmt.Println(avarage(12, 13, 1, 5, 78, 3))
	//we have data and pass it to the same function. look how we modify the argument
	fmt.Println(avarage(data...))

}

//variadic number of input parameters
func avarage(insf ...float64) float64 {
	fmt.Println(insf)
	fmt.Printf("%T \n", insf)
	var i int
	var sum float64 = 0
	for _, v := range insf {
		sum += v
		i++
	}
	//return sun/float64(len(insf))
	return sum / float64(i)
}

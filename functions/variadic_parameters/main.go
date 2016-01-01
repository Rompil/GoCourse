package main

import(
	"fmt"
)

func main(){
	fmt.Print("Среднее значение: ")
	fmt.Println(avarage(12,13,1,5,78,3))

}
//variadic number of input parameters
func avarage(insf ... float64) float64{
	fmt.Println(insf)
	fmt.Printf("%T \n", insf)
	var i int
	var sum float64 = 0
	for _,v := range insf{
		sum+=v
		i++
	}
	return sum/float64(i)
}
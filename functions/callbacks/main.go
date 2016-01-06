package main

import "fmt"

func main() {
	//we place anonymous function as a callback func
	fmt.Print(filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(inDig int) bool {
		if inDig%3 == 0 {
			return true
		}
		return false
	}))

}

func filter(insls []int, callback func(int) bool) []int {
	xs := []int{}
	for _, v := range insls {
		if callback(v) {
			xs = append(xs, v)
		}

	}
	return xs
}

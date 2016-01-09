package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	if p[i] < p[j] {
		return true
	}
	return false
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	//-------------------------------
	s := []string{"Zeno", "John", "Al", "Jenny"}
	//sort.StringSlice(s).Sort() //It is also works
	sl := sort.StringSlice(s)
	sort.Sort(sl)
	fmt.Println(sl)
	//-------------------------------
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	nsl := sort.IntSlice(n)
	sort.Sort(nsl)
	fmt.Println(nsl)
	//-------------------------------
	sort.Sort(sort.Reverse(nsl))
	fmt.Println(nsl)

}

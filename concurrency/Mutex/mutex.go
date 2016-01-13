package main

import(
	"fmt"
	"sync"
)

type AtomicInt struct{
	m sync.Mutex
	value int
}

func (s *AtomicInt) Add (a int){
	s.m.Lock()
	s.value+=a
	s.m.Unlock()
}

func (s AtomicInt) Value() int{
	s.m.Lock()
	v:= s.value
	s.m.Unlock()
	return v
}

func Race() {
	wait:= make(chan struct{})
	var v AtomicInt
	go func(){
		v.Add(1)
		close(wait)
	}()
	<-wait
	v.Add(1)
	fmt.Println(v.Value())
}

func main(){
	Race()
}

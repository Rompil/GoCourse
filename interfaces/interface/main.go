package main

import (
	"fmt"
	"math"
)

type square struct {
	Side float32
}
type circle struct {
	Radius float32
}

func (s square) area() float32 {
	return s.Side * s.Side
}

func (r circle) area() float32 {
	return math.Pi * r.Radius * r.Radius
}

type Shape interface {
	area() float32
}

func Info(s Shape) {
	fmt.Printf("A new %T's area is %v\n", s, s.area())
}
func main() {
	sq := square{4.0}
	ci := circle{5.0}
	Info(sq)
	Info(ci)
}

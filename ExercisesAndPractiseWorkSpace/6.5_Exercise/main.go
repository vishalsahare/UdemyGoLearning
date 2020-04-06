package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type square struct {
	length float32
	width  float32
}

func (c circle) area() float32 {
	return (math.Pi * (c.radius * 2))
}

func (s square) area() float32 {
	return s.length * s.width
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	oCircle := circle{
		radius: 2,
	}

	oSquare := square{
		length: 2,
		width:  2,
	}

	info(oCircle)
	info(oSquare)
}

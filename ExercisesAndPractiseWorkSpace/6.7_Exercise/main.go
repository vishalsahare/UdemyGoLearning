package main

import "fmt"

type square struct {
	length float32
	width  float32
}

func main() {

	oSquare := square{
		length: 2,
		width:  2,
	}

	vName := func() {
		fmt.Println("Area ", oSquare.length*oSquare.width)
	}

	vName()
}

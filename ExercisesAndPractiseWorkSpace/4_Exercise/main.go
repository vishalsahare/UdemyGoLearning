package main

import (
	"fmt"
)

type MyIntType int

var x MyIntType

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%v", x)
}

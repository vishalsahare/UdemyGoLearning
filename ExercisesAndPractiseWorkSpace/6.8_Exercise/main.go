package main

import "fmt"

func foo() func() int {
	return func() int {
		return 14
	}
}

func main() {
	vName := foo()
	fmt.Println(vName())
}

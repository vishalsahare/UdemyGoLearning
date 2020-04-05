package main

import (
	"fmt"
)

func main() {
	a := [5]int{40, 41, 42, 43, 44}

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", a)
}

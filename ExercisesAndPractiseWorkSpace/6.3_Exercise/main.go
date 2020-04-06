package main

import "fmt"

func foo(vardSlice ...int) int {
	total := 0
	for _, v := range vardSlice {
		total += v
	}
	return total
}

func bar(iSlice []int) int {
	total := 0
	for _, v := range iSlice {
		total += v
	}
	return total
}
func main() {
	a := []int{1, 2, 3, 4, 5}
	defer fmt.Println("Value from func foo()", foo(a...))
	fmt.Println("Value from func bar()", bar(a))
}

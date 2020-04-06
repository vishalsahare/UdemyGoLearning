package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func calculate(f func(x, y int) int) {
	fmt.Println(f(5, 5))

}

func main() {
	calculate(sum)
}

package main

import "fmt"

func foo() int {
	return 5
}

func boo() (int, string) {
	return 14, `Hello, World`
}

func main() {
	fmt.Println(foo())
	fmt.Println(boo())
}

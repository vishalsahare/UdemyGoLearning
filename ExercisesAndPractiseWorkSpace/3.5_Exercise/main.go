package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v i divided by 4, remainder: %d\n", i, i%4)
	}
}

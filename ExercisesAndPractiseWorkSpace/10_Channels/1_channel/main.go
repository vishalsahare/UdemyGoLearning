package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 14
	}()

	fmt.Println(<-c)
}

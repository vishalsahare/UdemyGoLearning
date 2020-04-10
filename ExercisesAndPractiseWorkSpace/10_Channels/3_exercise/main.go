//implement the range on the channel

package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	receive(c)

	fmt.Println("main: Done...")
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

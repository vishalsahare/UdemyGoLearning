//Check the comma idioms when fetch the value
//from channel

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 41
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

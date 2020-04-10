package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func func1() {
	fmt.Println("func1")
	wg.Done()
}

func func2() {
	fmt.Println("func2")
	wg.Done()
}

func main() {
	fmt.Println("Begin main")
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Go Routing\t", runtime.NumGoroutine())
	wg.Add(2)
	go func1()
	go func2()

	wg.Wait()
	fmt.Println("End main")
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0
var wg sync.WaitGroup
var mtx sync.Mutex

func inc(goRoutingNum int) {
	fmt.Println("GoRoutine no ", goRoutingNum)
	fmt.Println("Total routings ", runtime.NumGoroutine())

	mtx.Lock()
	temp := counter
	temp++
	counter = temp
	mtx.Unlock()
	wg.Done()
}

func main() {
	numOfGoRoutines := 100
	wg.Add(numOfGoRoutines)
	fmt.Println("CPUs ", runtime.NumCPU())
	for i := 0; i < numOfGoRoutines; i++ {
		go inc(i)
	}
	wg.Wait()
	fmt.Println("\nCounter ", counter)
}

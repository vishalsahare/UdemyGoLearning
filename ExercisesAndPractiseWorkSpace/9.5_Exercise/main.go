package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func inc(goRoutingNum int) {
	fmt.Println("GoRoutine no ", goRoutingNum)
	fmt.Println("Total routings ", runtime.NumGoroutine())

	atomic.AddInt64(&counter, 1)
	fmt.Println("Counter ", atomic.LoadInt64(&counter))

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

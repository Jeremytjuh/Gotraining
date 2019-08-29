package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var wg sync.WaitGroup

func main() {
	goroutines := 100
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go incrementer()
		fmt.Println("Counter:\t", counter)
		runtime.Gosched()
	}
	fmt.Println("Counter:\t", counter)
	wg.Wait()
}

func incrementer() {
	newvar := counter
	newvar++
	counter = newvar
	wg.Done()
}

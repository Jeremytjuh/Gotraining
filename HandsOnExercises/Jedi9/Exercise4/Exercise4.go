package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	goroutines := 100
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			mu.Lock()
			newvar := counter
			newvar++
			counter = newvar
			wg.Done()
			fmt.Println("Counter:\t", counter)
			mu.Unlock()
		}()
	}
	fmt.Println("Counter:\t", counter)
	wg.Wait()
}

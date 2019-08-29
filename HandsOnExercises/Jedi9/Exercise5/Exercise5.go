package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	goroutines := 100
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			mu.Lock()
			atomic.AddInt64(&counter, 1)
			wg.Done()
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			mu.Unlock()
		}()
	}
	fmt.Println("Counter:\t", counter)
	wg.Wait()
}

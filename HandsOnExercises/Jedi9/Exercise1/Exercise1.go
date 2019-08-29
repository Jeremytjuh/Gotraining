package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Printf("Number of Go routines before launching:\t\t %v\n", runtime.NumGoroutine())
	go hello1()
	fmt.Printf("Number of Go routines after launching 1:\t %v\n", runtime.NumGoroutine())
	go hello2()
	fmt.Printf("Number of Go routines after launching 2:\t %v\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("Number of Go routines after being done:\t\t %v\n", runtime.NumGoroutine())
}

func hello1() {
	fmt.Println("Hello1")
	wg.Done()
}

func hello2() {
	fmt.Println("Hello2")
	wg.Done()
}

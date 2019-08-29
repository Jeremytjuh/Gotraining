package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("Architecture:\t", runtime.GOARCH)
}

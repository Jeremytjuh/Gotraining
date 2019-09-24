package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filepath.Base(filename))
}

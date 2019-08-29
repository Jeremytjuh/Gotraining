package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	gorts := 10
	chans := 10
	for i := 0; i < gorts; i++ {
		go func() {
			for i2 := 0; i2 < chans; i2++ {
				c <- i2
			}
		}()
	}
	for i3 := 0; i3 < (gorts * chans); i3++ {
		fmt.Println(i3, <-c)
	}
}

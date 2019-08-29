package main

import (
	"fmt"
)

func main() {
	a := 6 //Choose what to compare with b
	b := 7 //Choose what to compare with a
	if a == b {
		fmt.Printf("The if statement returned a true value, because %v equals %v\n", a, b) //Choose what action to perform when the comparison of a and b equals true
	} else if a != b {
		fmt.Printf("The if statement returned a false value, because %v does not equal %v\n", a, b) //Choose what action to perform when the comparison of a and b equals false
	} else {
		fmt.Println("This message is nearly impossible to get")
	}
}

package main

import (
	"fmt"
)

func main() {
	a := make([]string, 50, 50)
	a = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}
	fmt.Printf("Full:\t\t%v\n", a)
	fmt.Printf("Length:\t\t%v\n", len(a))
	fmt.Printf("Capacity:\t%v\n", cap(a))
	fmt.Printf("Slice:\tNumber\tState\n")
	for i := 0; i < len(a); i++ {
		fmt.Printf("\t%v\t%v\n", i, a[i])
	}
}

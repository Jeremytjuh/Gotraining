package main

import (
	"fmt"
)

func main() {
	a := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	a["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for k, v := range a {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Printf("\t%v\t%v\n", i, v2)
		}
		fmt.Printf("\n")
	}
}

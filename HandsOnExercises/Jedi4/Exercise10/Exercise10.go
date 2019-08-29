package main

import (
	"fmt"
)

func main() {
	a := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	a["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	delete(a, "no_dr")

	for k, v := range a {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Printf("\t%v\t%v\n", i, v2)
		}
		fmt.Printf("\n")
	}
}

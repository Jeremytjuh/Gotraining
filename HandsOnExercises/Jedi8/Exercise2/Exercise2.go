package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	var people []person
	xb := []byte(s)
	err := json.Unmarshal(xb, &people)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range people {
		fmt.Printf("For: %v %v, age: %v\nThe sayings are:\n", v.First, v.Last, v.Age)
		for i2, v2 := range v.Sayings {
			fmt.Println(i2+1, v2)
		}
		fmt.Printf("\n")
	}
}

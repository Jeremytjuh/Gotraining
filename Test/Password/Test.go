package main

import (
	"log"

	"github.com/muesli/crunchy"
)

const pw = "Password12"

func main() {
	options := crunchy.Options{
		MinLength:      1,
		MinDiff:        5,
		MinDist:        3,
		DictionaryPath: "/~/../../usr/share/dict",
		CheckHIBP:      false,
	}
	validator := crunchy.NewValidatorWithOpts(options)

	err := validator.Check(pw)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

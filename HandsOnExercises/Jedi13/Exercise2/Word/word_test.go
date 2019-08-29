package word

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/02-code-finished/quote"
	"testing"
)

// func TestUseCount(t *testing.T) {

// }

func TestUseCount(t *testing.T) {
	s := UseCount("Yes Yes No Maybe")
	for k, v := range s {
		switch k {
		case "Yes":
			if v != 2 {
				t.Error("Expected", 2, "got", v)
			}
		case "No":
			if v != 1 {
				t.Error("Expected", 2, "got", v)
			}
		case "Maybe":
			if v != 1 {
				t.Error("Expected", 2, "got", v)
			}
		}
	}
}

func TestCount(t *testing.T) {
	s := Count("Hello everyone my name is Jeremy")
	if s != 6 {
		t.Error("Expected", 6, "Got", s)
	}
}

func ExampleCount() {
	fmt.Println(Count("Jeremy Nelemans"))
	// Output:
	// 2
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

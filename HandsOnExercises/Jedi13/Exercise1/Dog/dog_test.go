package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(10)
	if x != 70 {
		t.Error("Expected", 70, "Got", x)
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(10)
	if x != 70 {
		t.Error("Expected", 70, "Got", x)
	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	//Output:
	//21
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))
	//Output:
	//21
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}

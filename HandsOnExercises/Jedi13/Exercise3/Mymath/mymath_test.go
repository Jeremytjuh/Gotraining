package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 5},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 5
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

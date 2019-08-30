package mypack

import (
	"fmt"
	"testing"
)

func TestAddIntSlice(t *testing.T) {
	xi := []int{1, 2, 3, 4, 5}
	xii := []int{6, 7, 8, 9, 10}
	returnedvalue := AddIntSlice(xi, xii)
	answer := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if len(returnedvalue) == len(answer) {
		for i, v := range returnedvalue {
			if v != answer[i] {
				t.Error("Value in slices aren't the same")
			}
		}
	} else {
		t.Error("Length of slices aren't the same")
	}
}

func TestAddStringSlice(t *testing.T) {
	xs := []string{"A", "B", "C"}
	xss := []string{"D", "E", "F"}
	returnedvalue := AddStringSlice(xs, xss)
	answer := []string{"A", "B", "C", "D", "E", "F"}
	if len(returnedvalue) == len(answer) {
		for i, v := range returnedvalue {
			if v != answer[i] {
				t.Error("Value in slices aren't the same")
			}
		}
	} else {
		t.Error("Length of slices aren't the same")
	}
}

func TestSliceFloat64ToInt(t *testing.T) {
	xf := []float64{69.420, 420.69, 33.1, 23.53, 6767.66}
	returnedvalue := SliceFloat64ToInt(xf)
	answer := []int{69, 421, 33, 24, 6768}
	if len(returnedvalue) == len(answer) {
		for i, v := range returnedvalue {
			if v != answer[i] {
				t.Error("Value in slices aren't the same")
			}
		}
	} else {
		t.Error("Length of slices aren't the same")
	}
}

func TestSliceIntToFloat64(t *testing.T) {
	xf := []int{3, 5, 77, 9, 232, 43444343}
	returnedvalue := SliceIntToFloat64(xf)
	answer := []float64{3, 5, 77, 9, 232, 43444343}
	if len(returnedvalue) == len(answer) {
		for i, v := range returnedvalue {
			if v != answer[i] {
				t.Error("Value in slices aren't the same")
			}
		}
	} else {
		t.Error("Length of slices aren't the same")
	}
}

func TestCelToFahr(t *testing.T) {
	dc := 21
	returnedvalue := CelToFahr(float64(dc))
	answer := 70
	if returnedvalue != float64(answer) {
		t.Error("Expected", answer, "got", returnedvalue)
	}
}

func TestFahrToCel(t *testing.T) {
	df := 69
	returnedvalue := FahrToCel(float64(df))
	answer := 21
	if returnedvalue != float64(answer) {
		t.Error("Expected", answer, "got", returnedvalue)
	}
}

func TestConvertLitre(t *testing.T) {
	type test struct {
		litre float64
		metric string
		answer float64
	}

	tests := []test {
		test{10, "mm3", 10000000},
		test{10, "cm3", 10000},
		test{10, "dm3", 10},
		test{10, "m3", 0.01},
		test{10, "dam3", 0.00001},
		test{10, "hm3", 0.00000001},
		test{10, "km3", 0.00000000001},
		test{10, "youremad", 0},
	}

	for _, v := range tests {
		returnedvalue, err := ConvertLitre(v.litre, v.metric)
		if err != nil {
			fmt.Println("Value matches none of the cases")
		} else {
			if returnedvalue != v.answer {
				fmt.Println("Expected", v.answer, "got", returnedvalue)
			}
		}
	}
}

func ExampleAddIntSlice() {
	xi := []int{1, 2, 3, 4, 5}
	xii := []int{6, 7, 8, 9, 10}
	returnedvalue := AddIntSlice(xi, xii)
	fmt.Println(returnedvalue)
	//Output:
	//[1 2 3 4 5 6 7 8 9 10]
}

func ExampleAddStringSlice() {
	xs := []string{"A", "B", "C"}
	xss := []string{"D", "E", "F"}
	returnedvalue := AddStringSlice(xs, xss)
	fmt.Println(returnedvalue)
	//Output:
	//[A B C D E F]
}

func ExampleSliceFloat64ToInt() {
	xf := []float64{69.420, 420.69, 33.1, 23.53, 6767.66}
	returnedvalue := SliceFloat64ToInt(xf)
	fmt.Println(returnedvalue)
	//Output:
	//[69 421 33 24 6768]
}

func ExampleSliceIntToFloat64() {
	xi := []int{69, 420, 33, 23, 6767}
	returnedvalue := SliceIntToFloat64(xi)
	fmt.Println(returnedvalue)
	//Output:
	//[69 420 33 23 6767]
}

func ExampleCelToFahr() {
	dc := 21
	returnedvalue := CelToFahr(float64(dc))
	fmt.Println(returnedvalue)
	//Output:
	//70
}

func ExampleFahrToCel() {
	dc := 69
	returnedvalue := FahrToCel(float64(dc))
	fmt.Println(returnedvalue)
	//Output:
	//21
}

func ExampleConvertLitre() {
	litre := 10
	returnedvalue, err := ConvertLitre(float64(litre), "cm3")
	fmt.Println(returnedvalue, err)
	//Output:
	//10000 <nil>
}

func BenchmarkAddIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xi := []int{1, 2, 3, 4, 5}
		xii := []int{6, 7, 8, 9, 10}
		AddIntSlice(xi, xii)
	}
}

func BenchmarkAddStringSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xs := []string{"A", "B", "C"}
		xss := []string{"D", "E", "F"}
		AddStringSlice(xs, xss)
	}
}

func BenchmarkFloat64ToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xf := []float64{69.420, 420.69, 33.1, 23.53, 6767.66}
		SliceFloat64ToInt(xf)
	}
}

func BenchmarkIntToFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xi := []int{69, 420, 33, 23, 6767}
		SliceIntToFloat64(xi)
	}
}

func BenchmarkCelToFahr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dc float64 = 21
		CelToFahr(dc)
	}
}

func BenchmarkFahrToCel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var df float64 = 69
		FahrToCel(df)
	}
}

func BenchmarkConvertLitre(b *testing.B) {
	for i := 0;i < b.N; i++ {
		var litre float64 = 10
		ConvertLitre(litre, "d3")
	}
}

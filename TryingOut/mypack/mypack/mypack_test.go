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
		litre  float64
		metric string
		answer float64
	}

	tests := []test{
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

func TestConvM(t *testing.T) {
	type test struct {
		tocalc float64
		metric string
		answer []float64
	}

	tests := []test{
		test{10, "mm", []float64{10, 1, 0.1, 0.01, 0.001, 0.0001, 0.00001}},
		test{10, "cm", []float64{100, 10, 1, 0.1, 0.01, 0.001, 0.0001}},
		test{10, "dm", []float64{1000, 100, 10, 1, 0.1, 0.01, 0.001}},
		test{10, "m", []float64{10000, 1000, 100, 10, 1, 0.1, 0.01}},
		test{10, "dam", []float64{100000, 10000, 1000, 100, 10, 1, 0.1}},
		test{10, "hm", []float64{1000000, 100000, 10000, 1000, 100, 10, 1}},
		test{10, "km", []float64{10000000, 1000000, 100000, 10000, 1000, 100, 10}},
		test{10, "youremad", []float64{}},
	}

	for _, v := range tests {
		returnedvalue, err := ConvM(v.tocalc, v.metric)
		if err == nil {
			if len(returnedvalue) == len(v.answer) {
				for i, v := range returnedvalue {
					if v != returnedvalue[i] {
						t.Error("Value in slices aren't the same")
					}
				}
			} else {
				t.Error("Length of slices aren't the same")
			}
		} else {

		}
	}
}

func TestConvMArea(t *testing.T) {
	type test struct {
		tocalc float64
		metric string
		answer []float64
	}

	tests := []test{
		test{10, "mm2", []float64{10, 0.1, 0.001, 0.00001, 0.0000001, 0.000000001, 0.00000000001}},
		test{10, "cm2", []float64{1000, 10, 0.1, 0.001, 0.00001, 0.0000001, 0.000000001}},
		test{10, "dm2", []float64{100000, 1000, 10, 0.1, 0.001, 0.00001, 0.0000001}},
		test{10, "m2", []float64{10000000, 100000, 1000, 10, 0.1, 0.001, 0.00001}},
		test{10, "dam2", []float64{1000000000, 10000000, 100000, 1000, 10, 0.1, 0.001}},
		test{10, "hm2", []float64{100000000000, 1000000000, 10000000, 100000, 1000, 10, 0.1}},
		test{10, "km2", []float64{10000000000000, 100000000000, 1000000000, 10000000, 100000, 1000, 10}},
		test{10, "youremad", []float64{}},
	}

	for _, v := range tests {
		returnedvalue, err := ConvMArea(v.tocalc, v.metric)
		if err == nil {
			if len(returnedvalue) == len(v.answer) {
				for i, v := range returnedvalue {
					if v != returnedvalue[i] {
						t.Error("Value in slices aren't the same")
					}
				}
			} else {
				t.Error("Length of slices aren't the same")
			}
		} else {

		}
	}
}

func TestConvMCubic(t *testing.T) {
	type test struct {
		tocalc float64
		metric string
		answer []float64
	}

	tests := []test{
		test{10, "mm3", []float64{10, 0.01, 0.00001, 0.00000001, 0.00000000001, 0.00000000000001, 0.00000000000000001}},
		test{10, "cm3", []float64{10000, 10, 0.01, 0.00001, 0.00000001, 0.00000000001, 0.00000000000001}},
		test{10, "dm3", []float64{10000000, 10000, 10, 0.01, 0.00001, 0.00000001, 0.00000000001}},
		test{10, "m3", []float64{10000000000, 10000000, 10000, 10, 0.01, 0.00001, 0.00000001}},
		test{10, "dam3", []float64{10000000000000, 10000000000, 10000000, 10000, 10, 0.01, 0.00001}},
		test{10, "hm3", []float64{10000000000000000, 10000000000000, 10000000000, 10000000, 10000, 10, 0.01}},
		test{10, "km3", []float64{10000000000000000000, 10000000000000000, 10000000000000, 10000000000, 10000000, 10000, 10}},
		test{10, "youremad", []float64{}},
	}

	for _, v := range tests {
		returnedvalue, err := ConvMCubic(v.tocalc, v.metric)
		if err == nil {
			if len(returnedvalue) == len(v.answer) {
				for i, v := range returnedvalue {
					if v != returnedvalue[i] {
						t.Error("Value in slices aren't the same")
					}
				}
			} else {
				t.Error("Length of slices aren't the same")
			}
		} else {

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

func ExampleConvM() {
	var convertto float64 = 10
	returnedvalue, err := ConvM(convertto, "mm")
	fmt.Println(returnedvalue, err)
	//Output:
	//[10 1 0.1 0.01 0.001 0.0001 1e-05] <nil>
}

func ExampleConvMArea() {
	var convertto float64 = 10
	returnedvalue, err := ConvMArea(convertto, "m2")
	fmt.Println(returnedvalue, err)
	//Output:
	//[1e+07 100000 1000 10 0.1 0.001 1e-05] <nil>
}

func ExampleConvMCubic() {
	var convertto float64 = 10
	returnedvalue, err := ConvMCubic(convertto, "hm3")
	fmt.Println(returnedvalue, err)
	//Output:
	//[1e+16 1e+13 1e+10 1e+07 10000 10 0.01] <nil>
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
	for i := 0; i < b.N; i++ {
		var litre float64 = 10
		ConvertLitre(litre, "cm3")
	}
}

func BenchmarkConvM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tocalc float64 = 10
		ConvM(tocalc, "dm")
	}
}

func BenchmarkConvMArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tocalc float64 = 10
		ConvMArea(tocalc, "dam2")
	}
}

func BenchmarkConvMCubic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tocalc float64 = 10
		ConvMCubic(tocalc, "km3")
	}
}

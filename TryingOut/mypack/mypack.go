package mypack

//AddIntSlice adds a slice of int to another slice of int
func AddIntSlice(xaddto []int, xadd []int) []int {
	for i := 0; i < len(xadd); i++ {
		xaddto = append(xaddto, xadd[i])
	}
	return xaddto
}

//AddStringSlice adds a slice of string to another slice of string
func AddStringSlice(xaddto []string, xadd []string) []string {
	for i := 0; i < len(xadd); i++ {
		xaddto = append(xaddto, xadd[i])
	}
	return xaddto
}

//SliceIntToFloat64 converts a slice of type integer to a slice of type float64
//This one comes in handy when you want to calculate the area of a circle
func SliceIntToFloat64(xconvert []int) []float64 {
	xconverted := []float64{}
	for i := 0; i < len(xconvert); i++ {
		xconverted = append(xconverted, float64(xconvert[i]))
	}
	return xconverted
}

//SliceFloat64ToInt converts a slice of type float64 to a slice of type int
//During this process it removes any decimals necessary
func SliceFloat64ToInt(xconvert []float64) []int {
	xconverted := []int{}
	for i := 0; i < len(xconvert); i++ {
		xconverted = append(xconverted, int(xconvert[i]))
	}
	return xconverted
}

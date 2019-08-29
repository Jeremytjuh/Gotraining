package mypack

import (
	"math"
)

//AddIntSlice adds a slice of int to another slice of int
func AddIntSlice(xaddto, xadd []int) []int {
	for i := 0; i < len(xadd); i++ {
		xaddto = append(xaddto, xadd[i])
	}
	return xaddto
}

//AddStringSlice adds a slice of string to another slice of string
func AddStringSlice(xaddto, xadd []string) []string {
	for i := 0; i < len(xadd); i++ {
		xaddto = append(xaddto, xadd[i])
	}
	return xaddto
}

//SliceFloat64ToInt truncates a slice of type float64 to a slice of type int
//During this process it removes any decimals necessary
func SliceFloat64ToInt(xconvert []float64) []int {
	xconverted := []int{}
	for i := 0; i < len(xconvert); i++ {
		xconverted = append(xconverted, int(xconvert[i]))
	}
	return xconverted
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

// CelToFahr converts degrees Celsius to degrees Fahrenheit and rounds it to the nearest integer
func CelToFahr(cel float64) float64 {
	return math.Round((cel * 9 / 5) + 32)
}

//FahrToCel converts degrees Fahrenheit to degrees Celsius and rounds it to the nearest integer
func FahrToCel(fahr float64) float64 {
	return math.Round((fahr - 32) / 9 * 5)
}

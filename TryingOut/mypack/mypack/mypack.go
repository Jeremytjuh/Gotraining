package mypack

import (
	"fmt"
	"math"
)

//AddIntSlice adds a slice of int to another slice of int.
func AddIntSlice(xaddto, xadd []int) []int {
	for i := 0; i < len(xadd); i++ {
		xaddto = append(xaddto, xadd[i])
	}
	return xaddto
}

//AddStringSlice adds a slice of string to another slice of string.
func AddStringSlice(xaddto, xadd []string) []string {
	for i := 0; i < len(xadd); i++ {
		xaddto = append(xaddto, xadd[i])
	}
	return xaddto
}

//SliceFloat64ToInt truncates a slice of type float64 to a slice of type int.
//During this process it removes any decimals necessary.
func SliceFloat64ToInt(xconvert []float64) []int {
	xconverted := []int{}
	for i := 0; i < len(xconvert); i++ {
		xconverted = append(xconverted, int(math.Round(xconvert[i])))
	}
	return xconverted
}

//SliceIntToFloat64 converts a slice of type integer to a slice of type float64.
//This one comes in handy when you want to calculate the area of a circle.
func SliceIntToFloat64(xconvert []int) []float64 {
	xconverted := []float64{}
	for i := 0; i < len(xconvert); i++ {
		xconverted = append(xconverted, float64(xconvert[i]))
	}
	return xconverted
}

// CelToFahr converts degrees Celsius to degrees Fahrenheit and rounds it to the nearest integer.
func CelToFahr(cel float64) float64 {
	return math.Round((cel * 9 / 5) + 32)
}

//FahrToCel converts degrees Fahrenheit to degrees Celsius and rounds it to the nearest integer.
func FahrToCel(fahr float64) float64 {
	return math.Round((fahr - 32) / 9 * 5)
}

//ConvertLitre converts the given float64 into the unit stored in the string.
func ConvertLitre(litre float64, convertto string) (float64, error) {
	switch convertto {
	case "mm3":
		return litre * 1000000, nil
	case "cm3":
		return litre * 1000, nil
	case "dm3":
		return litre, nil
	case "m3":
		return litre / 1000, nil
	case "dam3":
		return litre / 1000000, nil
	case "hm3":
		return litre / 1000000000, nil
	case "km3":
		return litre / 1000000000000, nil
	default:
		return 0, nil
	}
}

//ConvM takes in a float64 and a string, the float is the number that it converts and the string represents the metric unit it is in.
//The float64 in this function is just the plain number
func ConvM(tocalc float64, metric string) ([]float64, error) {
	metricslice := []float64{}
	switch metric {
	case "mm":
		metricslice = []float64{
			tocalc,
			tocalc / 10,
			tocalc / 100,
			tocalc / 1000,
			tocalc / 10000,
			tocalc / 100000,
			tocalc / 1000000,
		}
		return metricslice, nil
	case "cm":
		metricslice = []float64{
			tocalc * 10,
			tocalc,
			tocalc / 10,
			tocalc / 100,
			tocalc / 1000,
			tocalc / 10000,
			tocalc / 100000,
		}
		return metricslice, nil
	case "dm":
		metricslice = []float64{
			tocalc * 100,
			tocalc * 10,
			tocalc,
			tocalc / 10,
			tocalc / 100,
			tocalc / 1000,
			tocalc / 10000,
		}
		return metricslice, nil
	case "m":
		metricslice = []float64{
			tocalc * 1000,
			tocalc * 100,
			tocalc * 10,
			tocalc,
			tocalc / 10,
			tocalc / 100,
			tocalc / 1000,
		}
		return metricslice, nil
	case "dam":
		metricslice = []float64{
			tocalc * 10000,
			tocalc * 1000,
			tocalc * 100,
			tocalc * 10,
			tocalc,
			tocalc / 10,
			tocalc / 100,
		}
		return metricslice, nil
	case "hm":
		metricslice = []float64{
			tocalc * 100000,
			tocalc * 10000,
			tocalc * 1000,
			tocalc * 100,
			tocalc * 10,
			tocalc,
			tocalc / 10,
		}
		return metricslice, nil
	case "km":
		metricslice = []float64{
			tocalc * 1000000,
			tocalc * 100000,
			tocalc * 10000,
			tocalc * 1000,
			tocalc * 100,
			tocalc * 10,
			tocalc,
		}
		return metricslice, nil
	default:
		return metricslice, fmt.Errorf("Error: unit matches none of the cases")
	}
}

//ConvMArea takes in a float64 and a string, the float is the number that it converts and the string represents the metric unit it is in.
//The float64 in this function is an area
func ConvMArea(tocalc float64, metric string) ([]float64, error) {
	metricslice := []float64{}
	switch metric {
	case "mm2":
		metricslice = []float64{
			tocalc,
			tocalc / 100,
			tocalc / 10000,
			tocalc / 1000000,
			tocalc / 100000000,
			tocalc / 10000000000,
			tocalc / 1000000000000,
		}
		return metricslice, nil
	case "cm2":
		metricslice = []float64{
			tocalc * 100,
			tocalc,
			tocalc / 100,
			tocalc / 10000,
			tocalc / 1000000,
			tocalc / 100000000,
			tocalc / 10000000000,
		}
		return metricslice, nil
	case "dm2":
		metricslice = []float64{
			tocalc * 10000,
			tocalc * 100,
			tocalc,
			tocalc / 100,
			tocalc / 10000,
			tocalc / 1000000,
			tocalc / 100000000,
		}
		return metricslice, nil
	case "m2":
		metricslice = []float64{
			tocalc * 1000000,
			tocalc * 10000,
			tocalc * 100,
			tocalc,
			tocalc / 100,
			tocalc / 10000,
			tocalc / 1000000,
		}
		return metricslice, nil
	case "dam2":
		metricslice = []float64{
			tocalc * 100000000,
			tocalc * 1000000,
			tocalc * 10000,
			tocalc * 100,
			tocalc,
			tocalc / 100,
			tocalc / 10000,
		}
		return metricslice, nil
	case "hm2":
		metricslice = []float64{
			tocalc * 100000000,
			tocalc * 1000000,
			tocalc * 10000,
			tocalc * 10000,
			tocalc * 100,
			tocalc,
			tocalc / 100,
		}
		return metricslice, nil
	case "km2":
		metricslice = []float64{
			tocalc * 1000000000000,
			tocalc * 10000000000,
			tocalc * 100000000,
			tocalc * 1000000,
			tocalc * 10000,
			tocalc * 100,
			tocalc,
		}
		return metricslice, nil
	default:
		return metricslice, fmt.Errorf("Error: unit matches none of the cases")
	}
}

//ConvMCubic takes in a float64 and a string, the float is the number that it converts and the string represents the metric unit it is in.
//The float64 in this function is cubic
func ConvMCubic(tocalc float64, metric string) ([]float64, error) {
	metricslice := []float64{}
	switch metric {
	case "mm3":
		metricslice = []float64{
			tocalc,
			tocalc / 1000,
			tocalc / 1000000,
			tocalc / 1000000000,
			tocalc / 1000000000000,
			tocalc / 1000000000000000,
			tocalc / 1000000000000000000,
		}
		return metricslice, nil
	case "cm3":
		metricslice = []float64{
			tocalc * 1000,
			tocalc,
			tocalc / 1000,
			tocalc / 1000000,
			tocalc / 1000000000,
			tocalc / 1000000000000,
			tocalc / 1000000000000000,
		}
		return metricslice, nil
	case "dm3":
		metricslice = []float64{
			tocalc * 1000000,
			tocalc * 1000,
			tocalc,
			tocalc / 1000,
			tocalc / 1000000,
			tocalc / 1000000000,
			tocalc / 1000000000000,
		}
		return metricslice, nil
	case "m3":
		metricslice = []float64{
			tocalc * 1000000000,
			tocalc * 1000000,
			tocalc * 1000,
			tocalc,
			tocalc / 1000,
			tocalc / 1000000,
			tocalc / 1000000000,
		}
		return metricslice, nil
	case "dam3":
		metricslice = []float64{
			tocalc * 1000000000000,
			tocalc * 1000000000,
			tocalc * 1000000,
			tocalc * 1000,
			tocalc,
			tocalc / 1000,
			tocalc / 1000000,
		}
		return metricslice, nil
	case "hm3":
		metricslice = []float64{
			tocalc * 1000000000000000,
			tocalc * 1000000000000,
			tocalc * 1000000000,
			tocalc * 1000000,
			tocalc * 1000,
			tocalc,
			tocalc / 1000,
		}
		return metricslice, nil
	case "km3":
		metricslice = []float64{
			tocalc * 1000000000000000000,
			tocalc * 1000000000000000,
			tocalc * 1000000000000,
			tocalc * 1000000000,
			tocalc * 1000000,
			tocalc * 1000,
			tocalc,
		}
		return metricslice, nil
	default:
		return metricslice, fmt.Errorf("Error: unit matches none of the cases")
	}
}

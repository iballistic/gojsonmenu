package gojsonmenu

import "math"

// Round rounds a float64 to the nearest specified decimal places.
func Round(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

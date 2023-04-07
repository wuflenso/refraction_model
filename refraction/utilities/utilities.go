package utilities

import "math"

func DegreeToRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func RadiansToDegree(deg float64) float64 {
	return deg * (180 / math.Pi)
}

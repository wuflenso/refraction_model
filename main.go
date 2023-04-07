package main

import (
	"fmt"

	"github.com/wuflenso/refraction_model/refraction"
	"github.com/wuflenso/refraction_model/refraction/utilities"
)

func main() {
	// I. Inputs
	// Convert angles from Degrees to Radians first
	velocities := []float64{900, 500, 200}
	layerThicknesses := []float64{500, 500, 500}
	grids := [][]float64{{0, 0}}
	angles := []float64{utilities.DegreeToRadians(20)}

	// II. Execute function
	resGrids, resAnglesRad := refraction.TraceRayRefraction(layerThicknesses, velocities, grids, angles)

	// III. Convert to angles to degrees
	var resAnglesDeg []float64
	for _, o := range resAnglesRad {
		resAnglesDeg = append(resAnglesDeg, utilities.RadiansToDegree(o))
	}

	// IV. Print results
	for i, _ := range resGrids {
		s := fmt.Sprintf("Coordinate: %.2f, θ2: %.2f°", resGrids[i], resAnglesDeg[i])
		fmt.Println(s)
	}
}

package refraction

import "math"

// Reference on Snells Law
// https://www.e-education.psu.edu/earth520/content/l4_p5.html

// This function returns the grid points and incidence angles exprienced by a single ray with an initial incidence angle and coordinates stated.
// Input Values:
// + layer thickness - any unit
// + layer velocities - any unit
// + coordinates (or grid points) - [x, y] | you have to define the first coordinate i.e [0, 0]
// + incidence angles - radians | you have to define the first angle incidence angle i.e 0.34
//
// Returns Values:
// + coordinates of each layer boundaries - [x, y]
// + incidence angles of each layer boundaries - radians
//
// Limitations:
// + Assuming the earth layer is horizontal
// + Attenuation not taken into consideration
// + May not be able to calculate large numbers
func TraceRayRefraction(layerThicknesses []float64, layerVelocities []float64, coordinates [][]float64, angles []float64) ([][]float64, []float64) {
	for i, _ := range layerVelocities {
		if i+1 == len(layerVelocities) {
			break
		}
		var o2 float64
		if i == 0 {
			o2 = calculateNextAngleOfIncidence(angles[0], layerVelocities[i], layerVelocities[i+1])
		} else {
			o2 = calculateNextAngleOfIncidence(angles[len(angles)-1], layerVelocities[i], layerVelocities[i+1])
		}
		coordinates = calculateNextGridPoints(layerThicknesses[i+1], o2, coordinates)
		angles = append(angles, o2)
	}

	return coordinates, angles
}

// Private Methods

// calculates the 'bottom' coordinate of a layer after being refracted
// the input angle is in radians
func calculateNextGridPoints(s float64, o float64, grids [][]float64) [][]float64 {
	startGrid := grids[len(grids)-1]
	endX := (math.Tan(o) * s) + startGrid[0]
	endY := startGrid[1] + s
	endGrid := []float64{endX, endY}
	return append(grids, endGrid)
}

// calculates the refracted angle
// The input angle is in radians
func calculateNextAngleOfIncidence(o1 float64, v1 float64, v2 float64) (o2 float64) {
	return math.Asin(v2 * math.Sin(o1) / v1)
}

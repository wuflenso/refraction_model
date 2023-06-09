package refraction

import (
	"math"
)

const (
	CriticalRefractionAngle = math.Pi / 2

	MessageAllRefracted       = "Refracted through all provided layers"
	MessageCriticalRefraction = "Critical angle refraction by the next layer"
	MessageTotalReflection    = "Total reflection by the next layer"
	MessageZeroIncidenceAngle = "Zero incidence angle"
)

// Reference on Snells Law
// https://www.e-education.psu.edu/earth520/content/l4_p5.html

// This function returns the coordinates and incidence angles exprienced by a single ray going through a multiple rock layers. The initial incidence angle and coordinates are required to be stated.
//
// Input Values:
// 	- layerThicknesses | Layer thickness in any unit
// 	- layerVelocities | Velocity of the layers in any unit
// 	- coordinates | Starting coordinates (or grid points) with [x, y] data structure for each. The first coordinate must be stated i.e [0, 0]
// 	- angles | Ray incidence angles in radians. The first incidence angle needs to be stated i.e 0.34
//
// Return Values:
// 	- coordinateList | Coordinates of each incidence layer boundaries with [x, y] data structure for each
// 	- angleList | Incidence angles of each incidence layer boundaries in radians
//	- message | A message indicating the result of the calculation
//
// Limitations:
// 	- Assuming the earth layers are horizontal
// 	- Attenuation not taken into consideration
// 	- May not be able to calculate large numbers
func TraceRayRefraction(layerThicknesses []float64, layerVelocities []float64, coordinates [][]float64, angles []float64) (coordinateList [][]float64, angleList []float64, message string) {
	for i := range layerVelocities {

		// Break if encounter last layer
		if i+1 == len(layerVelocities) {
			break
		}

		// Check if zero incidence angle
		if angles[0] == 0 {
			return coordinates, angles, MessageZeroIncidenceAngle
		}

		// Calculate angle of refraction/next angle of incidence
		var o2 float64
		if i == 0 {
			o2 = calculateNextAngleOfIncidence(angles[0], layerVelocities[i], layerVelocities[i+1])
		} else {
			o2 = calculateNextAngleOfIncidence(angles[len(angles)-1], layerVelocities[i], layerVelocities[i+1])
		}

		// 90 degrees refraction angle means it is critically refracted
		if math.Abs(o2) == CriticalRefractionAngle {
			return coordinates, angles, MessageCriticalRefraction
		}

		// NaN o2 means that the ray is totally reflected instead of refracted
		if math.IsNaN(o2) {
			return coordinates, angles, MessageTotalReflection
		}

		// Calculate next layer boundary coordinate and append angle of incidence
		coordinates = calculateNextGridPoints(layerThicknesses[i], o2, coordinates)
		angles = append(angles, o2)
	}

	return coordinates, angles, MessageAllRefracted
}

// Private Methods

// Calculates the 'bottom' coordinate of a layer after being refracted
// the input angle is in radians
func calculateNextGridPoints(s float64, o float64, grids [][]float64) [][]float64 {
	startGrid := grids[len(grids)-1]
	endX := (math.Tan(o) * s) + startGrid[0]
	endY := startGrid[1] + s
	endGrid := []float64{endX, endY}
	return append(grids, endGrid)
}

// Calculates the refracted angle
// The input angle is in radians
func calculateNextAngleOfIncidence(o1 float64, v1 float64, v2 float64) (o2 float64) {

	// round to 1 or -1 if outside asin input range of [-1, 1]
	asinInput := v2 * math.Sin(o1) / v1
	if asinInput > 1 || asinInput < -1 {
		asinInput = math.Round(asinInput)
	}
	return math.Asin(asinInput)
}

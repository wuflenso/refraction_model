package refraction_test

import (
	"math"
	"refraction_model/refraction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefraction_TraceRayRefraction(t *testing.T) {
	type args struct {
		layerThicknesses []float64
		layerVelocities  []float64
		coordinates      [][]float64
		angles           []float64
	}

	testCases := []struct {
		name                      string
		input                     args
		expectedOutputCoordinates [][]float64
		expectedOutputAngles      []float64
	}{
		{
			name: "normal case",
			input: args{
				layerThicknesses: []float64{100, 100},
				layerVelocities:  []float64{200, 500},
				coordinates:      [][]float64{{0, 0}},
				angles:           []float64{20 * (math.Pi / 180)},
			},
			expectedOutputCoordinates: [][]float64{{0, 0}, {164.89415762094316, 100}},
			expectedOutputAngles: []float64{
				0.3490658503988659,
				1.025647946453819},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			coordinates, angles := refraction.TraceRayRefraction(tc.input.layerThicknesses, tc.input.layerVelocities, tc.input.coordinates, tc.input.angles)
			assert.Equal(t, tc.expectedOutputCoordinates, coordinates)
			assert.Equal(t, tc.expectedOutputAngles, angles)
		})
	}
}

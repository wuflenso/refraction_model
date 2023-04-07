package utilities_test

import (
	"math"
	"refraction_model/refraction/utilities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilities_DegreeToRadians(t *testing.T) {
	testCases := []struct {
		name           string
		input          float64
		expectedOutput float64
	}{
		{
			name:           "normal case",
			input:          60,
			expectedOutput: 1.0471975511965976,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := utilities.DegreeToRadians(tc.input)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}

func TestUtilities_RadiansToDegree(t *testing.T) {
	testCases := []struct {
		name           string
		input          float64
		expectedOutput float64
	}{
		{
			name:           "normal case",
			input:          1.047,
			expectedOutput: 60,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := utilities.RadiansToDegree(tc.input)
			assert.Equal(t, tc.expectedOutput, math.Round(output))
		})
	}
}

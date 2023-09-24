package domain

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/package-counter/domain/calculation"
)

func TestCalculatePacks(t *testing.T) {
	// Define test cases with input and expected output
	testCases := []struct {
		input    float64
		expected []string
	}{
		{1, []string{"1x250"}},
		{251, []string{"1x500"}},
		// Add more test cases as needed
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %f", tc.input), func(t *testing.T) {
			result := calculation.CalculatePacks(tc.input)

			// Compare the result to the expected result
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}

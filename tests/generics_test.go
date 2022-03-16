package tests

import (
	"testing"

	go118 "github.com/ag9920/go-experiment/go_1_18"
	"github.com/stretchr/testify/assert"
)

func Test_generics(t *testing.T) {
	assert := assert.New(t)
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	expectedIntsSum := 46
	expectedFloatsSum := 62.97

	//Non-Generic Sums
	assert.EqualValues(expectedIntsSum, go118.SumInts(ints))
	assert.EqualValues(expectedFloatsSum, go118.SumFloats(floats))

	// Generic Sums
	assert.EqualValues(expectedIntsSum, go118.SumIntsOrFloats[string, int64](ints))
	assert.EqualValues(expectedFloatsSum, go118.SumIntsOrFloats[string, float64](floats))

	// Generic Sums, type parameters inferred
	// you can often omit the type arguments in the function call. Go can often infer them from your code.
	assert.EqualValues(expectedIntsSum, go118.SumIntsOrFloats(ints))
	assert.EqualValues(expectedFloatsSum, go118.SumIntsOrFloats(floats))

	// Generic Sums with Constraint
	assert.EqualValues(expectedIntsSum, go118.SumNumbers(ints))
	assert.EqualValues(expectedFloatsSum, go118.SumNumbers(floats))
}

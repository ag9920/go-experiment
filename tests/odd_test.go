package tests

import (
	"testing"

	"github.com/ag9920/go-experiment/bit"
)

func Test_OddEven(t *testing.T) {
	for _, c := range testCases {
		if bit.IsOdd(c.Number) != c.IsOdd {
			t.Errorf("check IsOdd failed, number=%v, isOdd=%v", c.Number, bit.IsOdd(c.Number))
		}
		if bit.IsEven(c.Number) == c.IsOdd {
			t.Errorf("check IsEven failed, number=%v, IsEven=%v", c.Number, bit.IsEven(c.Number))
		}
	}
}

var (
	testCases = []struct {
		Number int64
		IsOdd  bool
	}{
		{
			Number: 5,
			IsOdd:  true,
		},
		{
			Number: 524354342565567545,
			IsOdd:  true,
		},
		{
			Number: 8,
			IsOdd:  false,
		},
		{
			Number: 54359824,
			IsOdd:  false,
		},
		{
			Number: 0,
			IsOdd:  false,
		},
		{
			Number: -1,
			IsOdd:  true,
		},
	}
)

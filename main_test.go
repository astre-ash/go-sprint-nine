package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomeElements_SizesAndBounds(t *testing.T) {
	testCases := []struct {
		name      string
		inputSize int
		wantLen   int
	}{
		{
			name:      "Positive slice size",
			inputSize: 10,
			wantLen:   10,
		},
		{
			name:      "Single element size",
			inputSize: 1,
			wantLen:   1,
		},
		{
			name:      "Zero slice size (edge case)",
			inputSize: 0,
			wantLen:   0,
		},
		{
			name:      "Negative slice size (edge case)",
			inputSize: -6,
			wantLen:   0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := generateRandomElements(tc.inputSize)

			// Use require because if length is wrong, further loop makes no sense.
			require.Len(t, res, tc.wantLen)

			// Verify that numbers stay within the [minVal, maxVal] range.
			if tc.inputSize > 0 {
				for _, v := range res {
					assert.GreaterOrEqual(t, v, minVal, "Value %d should be >= %d", v, minVal)
					assert.Less(t, v, maxVal, "Value %d should be < %d", v, maxVal)
				}
			}
		})
	}
}
func TestMaximum(t *testing.T) {
	testCases := []struct {
		name    string
		input   []int
		wantMax int
	}{
		{
			name:    "Positive numbers",
			input:   []int{1, 5, 23, 7, 14},
			wantMax: 23,
		},
		{
			name:    "Negative numbers only",
			input:   []int{-10, -5, -23, -7},
			wantMax: -5,
		},
		{
			name:    "Mixed positive and negative numbers",
			input:   []int{-15, 0, 42, -1, 12},
			wantMax: 42,
		},
		{
			name:    "Identical numbers",
			input:   []int{7, 7, 7, 7},
			wantMax: 7,
		},
		{
			name:    "Single element",
			input:   []int{42},
			wantMax: 42,
		},
		{
			name:    "Empty slice (edge case)",
			input:   []int{},
			wantMax: 0,
		},
		{
			name:    "Nil slice (edge case)",
			input:   nil,
			wantMax: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := maximum(tc.input)
			assert.Equal(t, tc.wantMax, res)
		})
	}
}

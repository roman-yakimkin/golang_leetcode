package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyPow(t *testing.T) {
	testCases := []struct {
		x        float64
		n        int
		expected float64
	}{
		{
			x:        2,
			n:        2,
			expected: 4,
		},
		{
			x:        2,
			n:        0,
			expected: 1,
		},
		{
			x:        2,
			n:        10,
			expected: 1024,
		},
		{
			x:        2,
			n:        -2,
			expected: 0.25,
		},
		{
			x:        0,
			n:        -1,
			expected: 0,
		},
		{
			x:        1,
			n:        12345678,
			expected: 1,
		},
	}
	for i, testCase := range testCases {
		result := MyPow(testCase.x, testCase.n)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

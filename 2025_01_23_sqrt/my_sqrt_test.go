package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMySqrt(t *testing.T) {
	testCases := []struct {
		x        int
		expected int
	}{
		{
			x:        10,
			expected: 3,
		},
		{
			x:        4,
			expected: 2,
		},
		{
			x:        101,
			expected: 10,
		},
		{
			x:        121,
			expected: 11,
		},
	}

	for i, testCase := range testCases {
		result := MySqrt(testCase.x)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}

}

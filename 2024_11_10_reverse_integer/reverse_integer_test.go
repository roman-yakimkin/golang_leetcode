package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		x        int
		expected int
	}{
		{
			x:        123,
			expected: 321,
		},
		{
			x:        -123,
			expected: -321,
		},
		{
			x:        120,
			expected: 21,
		},
		{
			x:        1534236469,
			expected: 0,
		},
	}

	for i, testCase := range testCases {
		result := Reverse(testCase.x)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

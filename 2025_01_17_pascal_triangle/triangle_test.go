package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		numRows  int
		expected [][]int
	}{
		{
			numRows:  1,
			expected: [][]int{{1}},
		},
		{
			numRows:  2,
			expected: [][]int{{1}, {1, 1}},
		},
		{
			numRows:  5,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for i, testCase := range testCases {
		result := Generate(testCase.numRows)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}

}

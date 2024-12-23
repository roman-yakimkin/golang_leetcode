package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		src      [][]int
		expected [][]int
	}{
		{
			src:      [][]int{{4, 8}, {3, 6}},
			expected: [][]int{{3, 4}, {6, 8}},
		},
		{
			src:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			src:      [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			expected: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}
	for i, testCase := range testCases {
		result := Rotate(testCase.src)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

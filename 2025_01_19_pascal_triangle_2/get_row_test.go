package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRow(t *testing.T) {
	testCases := []struct {
		rowNum   int
		expected []int
	}{
		{
			rowNum:   0,
			expected: []int{1},
		},
		{
			rowNum:   1,
			expected: []int{1, 1},
		},
		{
			rowNum:   3,
			expected: []int{1, 3, 3, 1},
		},
		{
			rowNum:   4,
			expected: []int{1, 4, 6, 4, 1},
		},
	}

	for i, testCase := range testCases {
		result := GetRow(testCase.rowNum)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

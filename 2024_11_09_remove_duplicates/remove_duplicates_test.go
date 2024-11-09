package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{1, 1, 2},
			expected: 2,
		},
		{
			arr:      []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
		{
			arr:      []int{1},
			expected: 1,
		},
		{
			arr:      []int{1, 1, 2},
			expected: 2,
		},
	}

	for i, testCase := range testCases {
		result := RemoveDuplicates(testCase.arr)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

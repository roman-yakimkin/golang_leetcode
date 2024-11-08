package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		nums     []int
		val      int
		expected int
	}{
		{
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			nums:     []int{3, 3, 3, 3},
			val:      3,
			expected: 0,
		},
		{
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
		},
		{
			nums:     []int{0, 1, 0, 0, 0, 0, 0, 0, 1, 0},
			val:      0,
			expected: 2,
		},
		{
			nums:     []int{1},
			val:      1,
			expected: 0,
		},
		{
			nums:     []int{2},
			val:      3,
			expected: 1,
		},
	}
	for i, testCase := range testCases {
		result := RemoveElement(testCase.nums, testCase.val)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

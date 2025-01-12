package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			nums:     []int{1, 3},
			target:   3,
			expected: 1,
		},
		{
			nums:     []int{1, 3},
			target:   1,
			expected: 0,
		},
	}

	for i, testCase := range testCases {
		result := SearchInsert(testCase.nums, testCase.target)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}

}

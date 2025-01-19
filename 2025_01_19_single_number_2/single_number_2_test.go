package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, 3, 2},
			expected: 3,
		},
		{
			nums:     []int{0, 1, 0, 1, 0, 1, 99},
			expected: 99,
		},
	}

	for i, testCase := range testCases {
		result := SingleNumber(testCase.nums)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}

}

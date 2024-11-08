package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		digits   []int
		expected []int
	}{
		{
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			digits:   []int{0},
			expected: []int{1},
		},
		{
			digits:   []int{9},
			expected: []int{1, 0},
		},
		{
			digits:   []int{9, 9, 9, 9},
			expected: []int{1, 0, 0, 0, 0},
		},
	}
	for i, testCase := range testCases {
		result := PlusOne(testCase.digits)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		num      int
		expected bool
	}{
		{
			num:      12345,
			expected: false,
		},
		{
			num:      121,
			expected: true,
		},
		{
			num:      13331,
			expected: true,
		},
		{
			num:      -121,
			expected: false,
		},
		{
			num:      10,
			expected: false,
		},
	}
	for i, testCase := range testCases {
		result := IsPalindrome(testCase.num)
		assert.Equal(t, testCase.expected, result, "case #%d", i)
	}

}

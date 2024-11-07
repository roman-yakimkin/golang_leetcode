package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		strs     []string
		expected string
	}{
		{
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			strs:     []string{"aaaaaaaaa", "aaaaaaa", "aaaaaa"},
			expected: "aaaaaa",
		},
	}
	for i, testCase := range testCases {
		result := LongestCommonPrefix(testCase.strs)
		assert.Equal(t, testCase.expected, result, "case #%d", i)
	}
}

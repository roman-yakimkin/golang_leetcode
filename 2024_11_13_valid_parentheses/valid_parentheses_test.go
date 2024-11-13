package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		s        string
		expected bool
	}{
		{
			s:        "()",
			expected: true,
		},
		{
			s:        "()[]{}",
			expected: true,
		},
		{
			s:        "(}",
			expected: false,
		},
		{
			s:        "({[)}]",
			expected: false,
		},
		{
			s:        "([])",
			expected: true,
		},
		{
			s:        "[",
			expected: false,
		},
	}

	for i, testCase := range testCases {
		result := IsValid(testCase.s)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

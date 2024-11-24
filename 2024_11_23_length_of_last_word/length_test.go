package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		s        string
		expected int
	}{
		{
			s:        "Hello World",
			expected: 5,
		},
		{
			s:        "luffy is still joyboy",
			expected: 6,
		},
		{
			s:        "   fly me   to   the moon  ",
			expected: 4,
		},
	}

	for i, testCase := range testCases {
		result := LengthOfLastWord(testCase.s)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

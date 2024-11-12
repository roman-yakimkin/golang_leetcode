package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		a        string
		b        string
		expected string
	}{
		{
			a:        "1",
			b:        "1",
			expected: "10",
		},
		{
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
	}
	for i, testCase := range testCases {
		result := AddBinary(testCase.a, testCase.b)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		roman    string
		expected int
	}{
		{
			roman:    "I",
			expected: 1,
		},
		{
			roman:    "III",
			expected: 3,
		},
		{
			roman:    "IV",
			expected: 4,
		},
		{
			roman:    "LVIII",
			expected: 58,
		},
		{
			roman:    "LIX",
			expected: 59,
		},
		{
			roman:    "XC",
			expected: 90,
		},
		{
			roman:    "C",
			expected: 100,
		},
		{
			roman:    "ID",
			expected: 499,
		},
		{
			roman:    "MCMXCIV",
			expected: 1994,
		},
	}
	for i, testCase := range testCases {
		result := RomanToInt(testCase.roman)
		assert.Equal(t, testCase.expected, result, "case #%d", i+1)
	}
}

package main

func IsPalindrome(x int) bool {
	var digits []int
	if x < 0 {
		return false
	}
	for x > 0 {
		digits = append(digits, x%10)
		x = x / 10
	}
	l := len(digits)
	for i := 0; i < l/2; i++ {
		if digits[i] != digits[l-i-1] {
			return false
		}
	}
	return true
}

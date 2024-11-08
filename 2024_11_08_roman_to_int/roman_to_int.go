package main

func RomanToInt(s string) int {
	digits := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	for i := 0; i < len(s); i++ {
		current, next := digits[s[i]], 0
		if i < len(s)-1 {
			next = digits[s[i+1]]
		}
		if current >= next {
			result += current
		} else {
			result -= current
		}
	}

	return result
}

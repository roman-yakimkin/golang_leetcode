package main

func LengthOfLastWord(s string) int {
	var result int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			if result > 0 {
				return result
			}
		} else {
			result++
		}
	}

	return result
}

package main

func Reverse(x int) int {
	var result int

	for x != 0 {
		mod := x % 10

		minResult, maxResult := (-1<<31+mod)/10, (1<<31-1-mod)/10

		if result < minResult || result > maxResult {
			return 0
		}

		result = result*10 + mod
		x = x / 10
	}

	return result
}

package main

func PlusOne(digits []int) []int {
	result := make([]int, len(digits)+1)
	borrow := 1

	for i := len(result) - 1; i > 0; i-- {
		result[i] = digits[i-1] + borrow
		if result[i] == 10 {
			result[i] = 0
		} else {
			borrow = 0
		}
	}

	if borrow == 1 {
		result[0] = borrow
	} else {
		result = result[1:]
	}

	return result
}

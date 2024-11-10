package main

func MyPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if x == -1 && n%2 == 0 {
		return -x
	}
	if x == -1 || x == 1 {
		return x
	}

	result, i := x, 1
	for i != n {
		if n > 0 {
			result, i = result*x, i+1
		} else {
			newResult := result / x
			if newResult == result {
				return newResult
			}
			result, i = result/x, i-1
		}
	}

	return result
}

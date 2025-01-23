package main

func MySqrt(x int) int {
	l1, l2 := 0, 65535
	var result int

	for {
		result = (l1 + l2) >> 1

		v1 := result * result
		v2 := v1 + (result << 1) + 1
		if v1 <= x && v2 > x {
			break
		}

		if v1 > x {
			l2 = (l1 + l2) >> 1
		} else {
			l1 = (l1 + l2) >> 1
		}
	}
	return result
}

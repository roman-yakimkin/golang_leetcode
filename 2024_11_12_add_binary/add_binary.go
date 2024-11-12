package main

func AddBinary(a string, b string) string {
	cntA, cntB := len(a)-1, len(b)-1
	var borrow uint8
	var result string

	for cntA > -1 || cntB > -1 {
		var num1, num2 uint8
		if cntA > -1 {
			num1 = a[cntA] - 48
		}
		if cntB > -1 {
			num2 = b[cntB] - 48
		}
		sum := num1 + num2 + borrow
		borrow = 0
		if sum > 1 {
			borrow = 1
			sum = sum - 2
		}
		result = string(sum+48) + result

		cntA--
		cntB--
	}

	if borrow == 1 {
		result = "1" + result
	}

	return result
}

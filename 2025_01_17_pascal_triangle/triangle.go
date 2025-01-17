package main

func Generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)

		if i == 0 {
			result[i][0] = 1
		} else {
			for j := 0; j <= i; j++ {
				add1, add2 := 0, 0
				if j > 0 {
					add1 = result[i-1][j-1]
				}
				if j < i {
					add2 = result[i-1][j]
				}

				result[i][j] = add1 + add2
			}
		}
	}

	return result
}

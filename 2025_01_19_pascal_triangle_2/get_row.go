package main

func GetRow(rowIndex int) []int {
	triangle := make([][]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		triangle[i] = make([]int, i+1)

		if i == 0 {
			triangle[i][0] = 1
		} else {
			for j := 0; j <= i; j++ {
				add1, add2 := 0, 0
				if j > 0 {
					add1 = triangle[i-1][j-1]
				}
				if j < i {
					add2 = triangle[i-1][j]
				}

				triangle[i][j] = add1 + add2
			}
		}
	}

	return triangle[rowIndex]
}

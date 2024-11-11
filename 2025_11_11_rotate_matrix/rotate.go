package main

func Rotate(matrix [][]int) [][]int {
	size := len(matrix)

	for i := 0; i < size/2; i++ {
		for j := i; j < size-i-1; j++ {
			matrix[i][j], matrix[j][size-i-1], matrix[size-i-1][size-j-1], matrix[size-j-1][i] = matrix[size-j-1][i], matrix[i][j], matrix[j][size-i-1], matrix[size-i-1][size-j-1]
		}
	}

	return matrix
}

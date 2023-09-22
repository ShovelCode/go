package main

import (
	"fmt"
)

// getMinor computes the minor matrix for a given element of the matrix.
func getMinor(matrix [][]float64, row, col int) [][]float64 {
	size := len(matrix)
	minor := make([][]float64, size-1)
	for i := range minor {
		minor[i] = make([]float64, size-1)
	}

	for i := 0; i < size; i++ {
		if i == row {
			continue
		}
		for j := 0; j < size; j++ {
			if j == col {
				continue
			}

			r := i
			if i > row {
				r = i - 1
			}

			c := j
			if j > col {
				c = j - 1
			}

			minor[r][c] = matrix[i][j]
		}
	}
	return minor
}

// determinant recursively calculates the determinant of a matrix.
func determinant(matrix [][]float64) float64 {
	size := len(matrix)
	if size == 1 {
		return matrix[0][0]
	}
	if size == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	result := 0.0
	for j := 0; j < size; j++ {
		// Compute determinant for minors, and add/subtract from result
		sign := 1.0
		if j%2 != 0 {
			sign = -1.0
		}
		result += sign * matrix[0][j] * determinant(getMinor(matrix, 0, j))
	}
	return result
}

func main() {
	matrix := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Determinant:", determinant(matrix))
}

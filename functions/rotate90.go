package main

import (
	"fmt"
)

// rotate90 rotates the given matrix 90 degrees clockwise.
func rotate90(matrix [][]int) {
	n := len(matrix)

	// 1. Transpose the matrix.
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 2. Reverse the columns.
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Original Matrix:")
	printMatrix(matrix)

	rotate90(matrix)

	fmt.Println("\nRotated Matrix:")
	printMatrix(matrix)
}

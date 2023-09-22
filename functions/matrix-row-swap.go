func swapRows(matrix [][]int, row1, row2 int) {
	if row1 < 0 || row2 < 0 || row1 >= len(matrix) || row2 >= len(matrix) {
		fmt.Println("Invalid row indices")
		return
	}
	matrix[row1], matrix[row2] = matrix[row2], matrix[row1]
}

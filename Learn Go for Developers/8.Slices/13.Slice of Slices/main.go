package main

func createMatrix(rows, cols int) [][]int {

	matrix := make([][]int, 0)

	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {

			row = append(row, i*j)

		}
		matrix = append(matrix, row)
	}

	return matrix
}

package sprint

func TransposeMatrix(matrix [][]int) [][]int {

	rows := len(matrix)
	cols := len(matrix[0])

	transpose := make([][]int, cols)

	for i := range transpose {
		transpose[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transpose[j][i] = matrix[i][j]
		}
	}

	return transpose

}
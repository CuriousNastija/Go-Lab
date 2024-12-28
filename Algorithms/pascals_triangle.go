package sprint

func PascalsTriangle(n int) [][]int {

	triangle := [][]int{}

	if n == 0 {
		return triangle
	}

	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1

		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle = append(triangle, row)
	}

	return triangle

}
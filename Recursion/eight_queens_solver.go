package sprint

import (
	"strconv"
	"strings"
)

func isSafe(board [] int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board [i] == col || abs(board[i]-col) == abs(i-row) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func placeQuenns(board []int, row int, solutions *[]string) {
	if row == len(board) {
		solution := make([]string, len(board))
		for i, col := range board {
			solution[i] = strconv.Itoa(col + 1)
		}
		*solutions = append(*solutions, strings.Join(solution, ""))
		return
	}

	for col := 0; col < len(board); col++ {
		if isSafe(board, row, col) {
			board[row] = col
			placeQuenns(board, row+1, solutions)
		}
	}
}


func EightQueensSolver() string {
	board := make([]int, 8)
	var solutions []string
	placeQuenns(board, 0, &solutions)
	return strings.Join(solutions, "\n")

}
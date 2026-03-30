type Set = [9]int
func isValidSudoku(board [][]byte) bool {
	rows := make([]Set, 9)
	cols := make([]Set, 9)
	cube := make([]Set, 9)
	for i := range 9 {
		for j := range 9 {
			if board[i][j] == '.' {
				continue
			}
			idx := board[i][j] - '1'
			row := i
			col := j
			cub := i / 3 * 3 + j / 3
			if rows[row][idx] != 0 {
				return false
			}
			if cols[col][idx] != 0 {
				return false
			}
			if cube[cub][idx] != 0 {
				return false
			}
			rows[row][idx] = 1
			cols[col][idx] = 1
			cube[cub][idx] = 1
		}
	}
	return true
}

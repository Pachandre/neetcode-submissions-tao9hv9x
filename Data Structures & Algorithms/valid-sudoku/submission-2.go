func isValidSudoku(board [][]byte) bool {
 for i := 0; i < len(board); i++ {
  row := make(map[byte]struct{})
  column := make(map[byte]struct{})
  square := make(map[byte]struct{})
  for j := 0; j < len(board[i]); j++ {
   if _, ok := row[board[i][j]]; ok && board[i][j] != '.' {
    return false
   }
   if _, ok := column[board[j][i]]; ok && board[j][i] != '.' {
    return false
   }
   y := 3 * (i/3) + j/3
   x := (i%3)*3 + j%3
   if _, ok := square[board[y][x]]; ok && board[y][x] != '.' {
    return false
   }
   column[board[j][i]] = struct{}{}
   row[board[i][j]] = struct{}{}
   square[board[y][x]] = struct{}{}
  }
 }
 return true
}
При заполнении таблиц могут быть ограничения, стоит проверить есть ли они.
Например в задаче `*__36. Valid Sudoku__*` не нужно добавлять ячейки с '.'

```go
func isValidSudoku(board [][]byte) bool {
	row := make(map[[2]int]bool)
	col := make(map[[2]int]bool)
	box := make(map[[2]int]bool)

	for i := range len(board) {
		for j := range len(board[0]) {
			val := board[i][j]

			// это будет неверное решение, если здесь забыть проверку,
			// что ячейка содержит '.'. То есть вот эту проверку:
			// if val == '.' {
			//	continue
			// }

			iVal := int(val)
			boxNum := (i / 3) * 3 + (j / 3)

			if row[[2]int{i, iVal}] || col[[2]int{j, iVal}] || box[[2]int{boxNum, iVal}] {
				return false
			}

			row[[2]int{i, iVal}] = true
			col[[2]int{j, iVal}] = true
			box[[2]int{boxNum, iVal}] = true
		}
	}

	return true
}
```
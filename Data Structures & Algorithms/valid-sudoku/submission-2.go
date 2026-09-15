func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[int]struct{})
	columns := make(map[int]map[int]struct{})
	subgrids := make(map[[2]int]map[int]struct{})
	for r := range 9 {
		for c := range 9 {
			if board[r][c] == '.' {
				continue
			}
            if rows[r] == nil {
                rows[r] = make(map[int]struct{})
            }
            if columns[c] == nil {
                columns[c] = make(map[int]struct{})
            }
            subgrid := [2]int{r/3, c/3}
            if subgrids[subgrid] == nil {
                subgrids[subgrid] = make(map[int]struct{})
            }
            b, _ := strconv.Atoi(string(board[r][c]))
            _, rowExists := rows[r][b]
            _, columnExists := columns[c][b]
            _, subgridExists := subgrids[subgrid][b]
			if rowExists || columnExists || subgridExists {
				return false
			}
			rows[r][b] = struct{}{}
			columns[c][b] = struct{}{}
			subgrids[[2]int{r/3, c/3}][b] = struct{}{}
		}
	}
	return true
}

// https://leetcode.cn/problems/n-queens/
// h

package backtrack

import "strings"

func solveNQueens(n int) (res [][]string) {
	// 初始化棋盘，每个格子默认为 '.'
	board := make([][]byte, n)
	for i := range board {
		board[i] = []byte(strings.Repeat(".", n))
	}

	// 标记列是否被占用
	cols := make([]bool, n)
	// 标记对角线是否被占用（左上到右下：row + col）
	diag1 := make([]bool, 2*n-1)
	// 标记对角线是否被占用（右上到左下：n - 1 - row + col）
	diag2 := make([]bool, 2*n-1)

	var backtrack func(row int)

	// 按行遍历，因为每一行只能放一个皇后
	backtrack = func(row int) {
		// 如果已经放满 n 行，表示找到一个合法解
		if row == n {
			validBoard := make([]string, n)
			for i := 0; i < n; i++ {
				validBoard[i] = string(board[i])
			}
			res = append(res, validBoard)
			return
		}

		// 尝试在当前行的每一列放皇后
		for col := 0; col < n; col++ {
			d1 := row + col
			d2 := n - 1 - row + col

			// 剪枝：如果列或对角线冲突，则跳过
			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}

			// 做选择：放置皇后并标记状态
			board[row][col] = 'Q'
			cols[col], diag1[d1], diag2[d2] = true, true, true

			// 进入下一行
			backtrack(row + 1)

			// 撤销选择：回溯
			board[row][col] = '.'
			cols[col], diag1[d1], diag2[d2] = false, false, false
		}
	}

	backtrack(0)
	return
}

// 归档一个最开始写的错误的方法（over think 了）
/*
func solveNQueens_wrong(n int) (res [][]string) {
	var invalidPos = make(map[[]int]bool, n)
	var bt func(int, int)
	var countQ int
	var board = generateBoard(n)

	var addInvalidPos func([]int)
	var validStack [][]int
	addInvalidPos = func(newPos []int) {
		newInvalids := []int{}
		row, col := newPos[0], newPos[1]
		for pos, isValid := range invalidPos {
			if !isValid {
				continue
			}
			vrow, vcol := pos[0], pos[1]
			if vrow == row || vcol == col || vrow+vcol == row+col || vrow-vcol == row-col {
				invalidPos[pos] = true
				newInvalids = append(newInvalids, pos)
			}
		}
		validStack = append(validStack, newInvalids)
	}

	var removeInvalidPos func()
	removeInvalidPos = func() {
		invalids := validStack[len(validStack)-1]
		validStack = validStack[:len(validStack)-1]

		for pos := range invalids {
			invalidPos[pos] = false
		}
	}

	bt = func(row, col int) {
		if countQ == n {
			res = append(res, append([]string{}, board...))
			return
		}
		for r := row; r < n; r++ {
			row := []byte(board[r])
			for c := 0; c < n; c++ {
				row[c] = 'Q'
				board[r] = string(row)
				countQ++
				addInvalidPos([]int{r, c})

				bt()

				row[c] = '.'
				board[r] = string(row)
				countQ--
				removeInvalidPos()
			}
		}
	}
	bt(0, 0)
	return
}

func generateBoard(n int) []string {
	row := strings.Repeat(".", n) // 一行全是 '.'
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = row
	}
	return board
}
*/

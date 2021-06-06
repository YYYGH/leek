package backtrack

import "fmt"

/*
51. N 皇后
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

var nqueens [][]string

func SolveNQueens(n int) [][]string {
	// 构造棋盘
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = '.'
	}
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = append([]byte{}, s...)
	}
	nqueens = make([][]string, 0)
	backtrack(board, 0)
	return nqueens
}

func backtrack(board [][]byte, row int) {
	if row == len(board) {
		nqueens = append(nqueens, []string{})
		fmt.Println(len(nqueens))
		for _, v := range board {
			fmt.Println(len(nqueens), v)
			nqueens[len(nqueens)-1] = append(nqueens[len(nqueens)-1], string(v))
		}
		// res = append(res, append([]string{}, board...))
		return
	}

	for col := 0; col < len(board); col++ {
		if !isValid(board, row, col) {
			continue
		}
		// 可以放置皇后
		board[row][col] = 'Q'
		// 寻找下一行可放置皇后位置
		backtrack(board, row+1)
		// 还原棋盘
		board[row][col] = '.'
	}
}

func isValid(board [][]byte, row, col int) bool {
	j := col - 1
	// 判断左上方是否有皇后
	for i := row - 1; i >= 0 && j >= 0; i-- {
		if board[i][j] == 'Q' {
			return false
		}
		j--
	}
	// 判断右上方是否有皇后
	i := row - 1
	for j := col + 1; j < len(board) && i >= 0; j++ {
		if board[i][j] == 'Q' {
			return false
		}
		i--
	}

	// 判断列上是否有皇后
	for i := row; i >= 0; i-- {
		if board[i][col] == 'Q' {
			return false
		}
	}

	return true
}

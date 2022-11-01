package problem100

/*
*
130. 被围绕的区域
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
*/
var n, m int

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m = len(board), len(board[0])
	for i := 0; i < n; i++ {
		dfsSolve(board, i, 0)
		dfsSolve(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		dfsSolve(board, 0, i)
		dfsSolve(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfsSolve(board [][]byte, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfsSolve(board, x+1, y)
	dfsSolve(board, x-1, y)
	dfsSolve(board, x, y+1)
	dfsSolve(board, x, y-1)
}

package important

//37. 解数独
//编写一个程序，通过填充空格来解决数独问题。
//
//数独的解法需 遵循如下规则：
//
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//数独部分空格内已填入了数字，空白格用 '.' 表示。

//提示：
//
//board.length == 9
//board[i].length == 9
//board[i][j] 是一位数字或者 '.'
//题目数据 保证 输入数独仅有一个解

// 题解：参考“代码随想录”解题思路
func solveSudoku(board [][]byte) {
	backtrack(board)
}

func backtrack(board [][]byte) bool {
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				for val := '1'; val <= '9'; val++ {
					if isValid(board, i, j, byte(val)) {
						board[i][j] = byte(val)
						if backtrack(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, val byte) bool {
	for r := 0; r < 9; r++ {
		if board[r][col] == val {
			return false
		}
	}

	for c := 0; c < 9; c++ {
		if board[row][c] == val {
			return false
		}
	}

	startRidx := row / 3 * 3
	startCidx := col / 3 * 3
	for r := startRidx; r < startRidx+3; r++ {
		for c := startCidx; c < startCidx+3; c++ {
			if board[r][c] == val {
				return false
			}
		}
	}

	return true
}

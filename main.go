package main

import (
	"fmt"
)

func main() {
	//[[1,4,5],[1,3,4],[2,6]]
	//l1 := NewListNodeS(1, 2, 3, 4, 5)
	//val := []*ListNode{NewListNodeS(1, 4, 5), NewListNodeS(1, 3, 4), NewListNodeS(2, 6)}
	//val := []int{1, 3, 5, 7, 9, 11, 13}

	val := [][]string{
		{".", ".", ".", ".", "5", ".", ".", "1", "."},
		{".", "4", ".", "3", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "3", ".", ".", "1"},
		{"8", ".", ".", ".", ".", ".", ".", "2", "."},
		{".", ".", "2", ".", "7", ".", ".", ".", "."},
		{".", "1", "5", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "2", ".", ".", "."},
		{".", "2", ".", "9", ".", ".", ".", ".", "."},
		{".", ".", "4", ".", ".", ".", ".", ".", "."},
	}
	v := isValidSudoku(val)
	//v := reverseList(l1)
	fmt.Println(v)
}

//32. 下一个排列
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//示例 1：
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//示例 2：
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//示例 3：
//输入：s = ""
//输出：0
//
//提示：
//0 <= s.length <= 3 * 104
//s[i] 为 '(' 或 ')'

//题解：使用栈的思想方法
func isValidSudoku(board [][]byte) bool {
	var rows, cols [9][9]int
	var box [3][3][9]int

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}

			val := board[row][col] - '1'
			rows[row][val]++
			cols[col][val]++
			box[row/3][col/3][val]++
			if rows[row][val] > 1 || cols[col][val] > 1 || box[row/3][col/3][val] > 1 {
				return false
			}
		}
	}

	return true
}

func getKey(row int, col int) int {
	return row/3*10 + col/3
}

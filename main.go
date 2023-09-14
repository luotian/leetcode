package main

import (
	"fmt"
)

func main() {
	//[[1,4,5],[1,3,4],[2,6]]
	//l1 := NewListNodeS(1, 2, 3, 4, 5)
	//val := []*ListNode{NewListNodeS(1, 4, 5), NewListNodeS(1, 3, 4), NewListNodeS(2, 6)}
	val := []int{4, 5, 6, 7, 0, 1, 2}
	v := search(val, 0)
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
func search(nums []int, target int) int {
	if len(nums) <= 2 {
		for i, v := range nums {
			if v == target {
				return i
			}
		}

		return -1
	}

	l := len(nums)
	lIdx := 0
	rIdx := l - 1
	lVal := nums[lIdx]
	rVal := nums[rIdx]

	switch target {
	case nums[lIdx]:
		return lIdx
	case nums[rIdx]:
		return rIdx
	}

	var mid int
	for lIdx <= rIdx {
		mid = (lIdx + rIdx) / 2
		if nums[mid] == target {
			return mid
		}

		if lVal <= nums[mid] {
			if lVal <= target && target <= nums[mid] {
				rIdx = mid - 1
			} else {
				lIdx = mid + 1
			}
		} else {
			if rVal >= target && target >= nums[mid] {
				lIdx = mid + 1
			} else {
				rIdx = mid - 1
			}
		}
	}

	return -1
}

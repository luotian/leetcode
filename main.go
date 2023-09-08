package main

import (
	"fmt"
)

func main() {
	//[[1,4,5],[1,3,4],[2,6]]
	//l1 := NewListNodeS(1, 2, 3, 4, 5)
	//val := []*ListNode{NewListNodeS(1, 4, 5), NewListNodeS(1, 3, 4), NewListNodeS(2, 6)}
	//val := []int{3, 2, 2, 3}
	v := strStr("aabaaabaaac", "aabaaac")
	//v := reverseList(l1)
	fmt.Println(v)
}

//No.25
//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			fIdx := i
			cur := i
			var newI = -2

			cur++
			for cur < len(haystack) && cur-fIdx < len(needle) && haystack[cur] == needle[cur-fIdx] {
				if newI == -2 && haystack[cur] == needle[0] {
					newI = cur - 1
				}
				cur++
			}

			if cur-fIdx == len(needle) {
				return fIdx
			}

			if newI == -2 {
				newI = cur - 1
			}
			i = newI
		}
	}

	return -1
}

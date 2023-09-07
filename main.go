package main

import (
	"fmt"
)

func main() {
	//[[1,4,5],[1,3,4],[2,6]]

	l1 := NewListNodeS(1, 2, 3, 4, 5)

	//val := []*ListNode{NewListNodeS(1, 4, 5), NewListNodeS(1, 3, 4), NewListNodeS(2, 6)}

	v := swapPairs(l1)
	fmt.Println(v)
}

//给你一个链表数组，每个链表都已经按升序排列。
//
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//示例 1：
//
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//示例 2：
//
//输入：lists = []
//输出：[]
//示例 3：
//
//输入：lists = [[]]
//输出：[]
//
//
//提示：
//
//k == lists.length
//0 <= k <= 10^4
//0 <= lists[i].length <= 500
//-10^4 <= lists[i][j] <= 10^4
//lists[i] 按 升序 排列
//lists[i].length 的总和不超过 10^4

func swapPairs(head *ListNode) *ListNode {
	var cursorNode = &ListNode{}
	var headNode = cursorNode

	cursorNode.Next = head
	for cursorNode.Next != nil && cursorNode.Next.Next != nil {
		first := cursorNode.Next
		second := cursorNode.Next.Next
		first.Next = second.Next
		second.Next = first
		cursorNode.Next = second
		cursorNode = cursorNode.Next.Next
	}

	return headNode.Next
}

package main

import (
	"fmt"
)

func main() {
	//[[1,4,5],[1,3,4],[2,6]]

	l1 := NewListNodeS(1, 2, 3, 4, 5)

	//val := []*ListNode{NewListNodeS(1, 4, 5), NewListNodeS(1, 3, 4), NewListNodeS(2, 6)}

	v := reverseKGroup(l1, 3)
	
	fmt.Println(v)
}

//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}

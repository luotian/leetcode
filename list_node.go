package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func NewListNodeS(val ...int) *ListNode {
	n := len(val)
	var lastNode *ListNode
	for i := n - 1; i >= 0; i-- {
		cur := NewListNode(val[i])
		if lastNode != nil {
			cur.Next = lastNode
		}
		lastNode = cur
	}

	return lastNode
}

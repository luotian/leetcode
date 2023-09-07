package important

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz
//Definition for singly-linked list.

//v1 := &ListNode{
//Val: 1,
//}
//v2 := &ListNode{
//Val: 2,
//}
//v3 := &ListNode{
//Val: 4,
//}
//v1.Next = v2
//v2.Next = v3
//vv1 := &ListNode{
//Val: 1,
//}
//vv2 := &ListNode{
//Val: 3,
//}
//vv3 := &ListNode{
//Val: 4,
//}
//vv1.Next = vv2
//vv2.Next = vv3

type ListNode struct {
	Val  int
	Next *ListNode
}

// input: [1,2,3,4,5] n=2
// output:[1,2,3,5]
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first, second := head, head
	for i := 0; i < n; i++ {
		first = first.Next
		if first == nil {
			return second.Next
		}
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return head
}

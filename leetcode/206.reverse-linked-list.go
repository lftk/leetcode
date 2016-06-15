package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	/*
		node := reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
		return node
	*/
	node, next := head, head.Next
	for next != nil {
		tmp := next.Next
		next.Next = node
		node = next
		next = tmp
	}
	head.Next = nil
	return node
}

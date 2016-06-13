package leetcode

import "testing"

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

func TestReverseList(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n := reverseList(n1)
	if n.Val != 5 {
		t.Fatal("reverseList failed")
	}
}

package leetcode

import "testing"

func Test_removeNthFromEnd(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n := removeNthFromEnd(n1, 5)
	if n == nil || n.Val != 2 {
		t.Fatal("removeNthFromEnd failed")
	}
}

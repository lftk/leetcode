package leetcode

import "testing"

func Test_isPalindromeLinkedList(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 2}
	n5 := &ListNode{Val: 1}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	b1 := isPalindromeLinkedList(n1)
	if !b1 {
		t.Fatal("12321 is palindrome")
	}

	n6 := &ListNode{Val: 1}
	n7 := &ListNode{Val: 2}
	n8 := &ListNode{Val: 2}
	n9 := &ListNode{Val: 1}
	n6.Next = n7
	n7.Next = n8
	n8.Next = n9
	b2 := isPalindromeLinkedList(n6)
	if !b2 {
		t.Fatal("1212 is palindrome")
	}
}

package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n1
	b := hasCycle(n1)
	if b == false {
		t.Fatal(b)
	}
}

package leetcode

import (
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 2}
	l4 := &ListNode{Val: 3}
	l5 := &ListNode{Val: 3}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l := deleteDuplicates(l1)
	vals := []int{1, 2, 3}
	var i int
	b := true
	for n := l; n != nil; n = n.Next {
		if i >= len(vals) || vals[i] != n.Val {
			b = false
			break
		}
		i++
	}
	if b == false {
		t.Error("1->1->2->3->3")
	}
}

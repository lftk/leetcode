package leetcode

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l6 := &ListNode{Val: 6}
	l7 := &ListNode{Val: 7}
	l8 := &ListNode{Val: 8}
	l1.Next = l3
	l3.Next = l5
	l5.Next = l7
	l2.Next = l4
	l4.Next = l6
	l6.Next = l8
	l := mergeTwoLists(l1, l2)
	var vals []int
	for n := l; n != nil; n = n.Next {
		vals = append(vals, n.Val)
	}
	for i, v := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		if i >= len(vals) || v != vals[i] {
			t.Error(i, v, vals[i])
			break
		}
	}
}

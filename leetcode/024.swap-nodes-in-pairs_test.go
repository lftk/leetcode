package leetcode

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l := swapPairs(l1)
	var vals []int
	for n := l; n != nil; n = n.Next {
		vals = append(vals, n.Val)
	}
	for i, v := range []int{2, 1, 4, 3} {
		if i >= len(vals) || v != vals[i] {
			t.Error(i, v, vals[i])
			break
		}
	}
}

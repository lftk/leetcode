package leetcode

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l1.Next = l2
	l2.Next = l3
	lst := removeElements(l1, 2)
	nums := []int{1, 3}
	for i, l := 0, lst; l != nil; i, l = i+1, l.Next {
		if i >= len(nums) || nums[i] != l.Val {
			t.Error(i, l.Val)
		}
	}
}

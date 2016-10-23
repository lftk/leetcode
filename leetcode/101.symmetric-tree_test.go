package leetcode

import (
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	r := &TreeNode{Val: 1}
	l1 := &TreeNode{Val: 2}
	r1 := &TreeNode{Val: 2}
	l2 := &TreeNode{Val: 3}
	r2 := &TreeNode{Val: 3}
	r.Left = l1
	r.Right = r1
	l1.Left = l2
	r1.Right = r2
	b1 := isSymmetric(r)
	if !b1 {
		t.Error(b1, "true")
	}
}

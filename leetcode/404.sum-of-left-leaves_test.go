package leetcode

import (
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n := sumOfLeftLeaves(n1)
	if n != 6 {
		t.Error(n, 6)
	}
}

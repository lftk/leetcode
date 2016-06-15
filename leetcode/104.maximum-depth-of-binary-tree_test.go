package leetcode

import "testing"

func Test_maxDepth(t *testing.T) {
	n := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	if v := maxDepth(n); v != 2 {
		t.Fatalf("%d != 2", v)
	}
}

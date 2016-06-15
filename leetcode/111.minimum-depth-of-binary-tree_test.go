package leetcode

import "testing"

func Test_minDepth(t *testing.T) {
	n := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	if v := minDepth(n); v != 2 {
		t.Fatalf("%d != 2", v)
	}
}

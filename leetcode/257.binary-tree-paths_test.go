package leetcode

import "testing"

func Test_binaryTreePaths(t *testing.T) {
	n := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	s := binaryTreePaths(n)
	if len(s) != 2 {
		t.Fatal("len(s) != 2")
	}
	if s[0] != "1->2" || s[1] != "1->3" {
		t.Fatal(s)
	}
}

package leetcode

import "testing"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var l, r int
	if root.Left != nil {
		l = maxDepth(root.Left)
	}
	if root.Right != nil {
		r = maxDepth(root.Right)
	}
	if l > r {
		return l + 1
	}
	return r + 1
}

func TestMaxDepth(t *testing.T) {
	n := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	if v := maxDepth(n); v != 2 {
		t.Fatalf("%d != 2", v)
	}
}

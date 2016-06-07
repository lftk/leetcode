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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
        return 1
    }
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }
    l, r := minDepth(root.Left), minDepth(root.Right)
    if l < r {
        return l + 1
    }
    return r + 1
}

func TestMinDepth(t *testing.T) {
	n := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
	}
	if v := minDepth(n); v != 2 {
		t.Fatalf("%d != 2", v)
	}
}

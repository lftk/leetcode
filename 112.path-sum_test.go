package leetcode

import (
	"testing"
)

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum -= root.Val; root.Left == nil && root.Right == nil {
		return sum == 0
	}
	if root.Left != nil && hasPathSum(root.Left, sum) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, sum) {
		return true
	}
	return false
}

func TestHasPathSum(t *testing.T) {
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}
	n8 := &TreeNode{Val: 8}
	n11 := &TreeNode{Val: 11}
	n13 := &TreeNode{Val: 13}
	n3 := &TreeNode{Val: 3}
	n7 := &TreeNode{Val: 7}
	n2 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1}
	n5.Left = n4
	n5.Right = n8
	n4.Left = n11
	n8.Left = n13
	n8.Right = n3
	n11.Left = n7
	n11.Right = n2
	n3.Right = n1
	if !hasPathSum(n5, 22) {
		t.Fatal("no find 5->4->11->2")
	}
}

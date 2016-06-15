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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2
	root := &TreeNode{Val: nums[m]}
	if 0 < m {
		root.Left = sortedArrayToBST(nums[:m])
	}
	if m < len(nums)-1 {
		root.Right = sortedArrayToBST(nums[m+1:])
	}
	return root
}

func TestSortedArrayToBST(t *testing.T) {
	root := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6})
	if root.Val != 4 {
		t.Fatal(root)
	}
	if root.Left.Val != 2 {
		t.Fatal(root)
	}
}

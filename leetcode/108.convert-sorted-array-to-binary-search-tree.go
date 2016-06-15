package leetcode

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

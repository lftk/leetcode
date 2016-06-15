package leetcode

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	var ss []string
	val := strconv.Itoa(root.Val) + "->"
	if root.Left != nil {
		for _, s := range binaryTreePaths(root.Left) {
			ss = append(ss, val+s)
		}
	}
	if root.Right != nil {
		for _, s := range binaryTreePaths(root.Right) {
			ss = append(ss, val+s)
		}
	}
	return ss
}

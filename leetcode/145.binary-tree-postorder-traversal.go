package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	if left := postorderTraversal(root.Left); left != nil {
		ret = append(ret, left...)
	}
	if right := postorderTraversal(root.Right); right != nil {
		ret = append(ret, right...)
	}
	ret = append(ret, root.Val)

	return ret
}

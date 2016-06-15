package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	for child := []*TreeNode{root}; child != nil; child = levelSubTreeNode(child) {
		vals := make([]int, len(child))
		for i, n := range child {
			vals[i] = n.Val
		}
		ret = append(ret, vals)
	}
	return ret
}

func levelSubTreeNode(nodes []*TreeNode) (child []*TreeNode) {
	for _, n := range nodes {
		if n.Left != nil {
			child = append(child, n.Left)
		}
		if n.Right != nil {
			child = append(child, n.Right)
		}
	}
	return
}

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := append([]int{}, root.Val)
	if left := preorderTraversal(root.Left); left != nil {
		ret = append(ret, left...)
	}
	if right := preorderTraversal(root.Right); right != nil {
		ret = append(ret, right...)
	}

	return ret
}

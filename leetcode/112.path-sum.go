package leetcode

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

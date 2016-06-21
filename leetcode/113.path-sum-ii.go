// https://leetcode.com/problems/path-sum-ii/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	if sum -= root.Val; root.Left == nil && root.Right == nil {
		if sum == 0 {
			return [][]int{{root.Val}}
		}
		return nil
	}

	var ret [][]int
	if root.Left != nil {
		ret = append(ret, pathSum(root.Left, sum)...)
	}
	if root.Right != nil {
		ret = append(ret, pathSum(root.Right, sum)...)
	}
	for i := 0; i < len(ret); i++ {
		ret[i] = append([]int{root.Val}, ret[i]...)
	}
	return ret
}

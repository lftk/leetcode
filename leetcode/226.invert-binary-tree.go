// https://leetcode.com/problems/invert-binary-tree/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root != nil {
        root.Left, root.Right = root.Right, root.Left
        invertTree(root.Left)
        invertTree(root.Right)
    }
    return root
}

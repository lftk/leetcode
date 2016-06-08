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

func TestPostorderTraversal(t *testing.T) {
	t1 := newTreeNode(1)
	t2 := newTreeNode(2)
	t3 := newTreeNode(3)

	t1.setRight(t2)
	t2.setLeft(t3)

	vals := []int{3, 2, 1}
	for i, v := range postorderTraversal(t1) {
		if vals[i] != v {
			t.Fatalf("%d %d %d", i, v, vals[i])
		}
	}
}

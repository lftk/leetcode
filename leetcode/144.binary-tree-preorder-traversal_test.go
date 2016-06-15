package leetcode

import "testing"

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func (t *TreeNode) setLeft(left *TreeNode) {
	t.Left = left
}

func (t *TreeNode) setRight(right *TreeNode) {
	t.Right = right
}

func Test_preorderTraversal(t *testing.T) {
	t1 := newTreeNode(1)
	t2 := newTreeNode(2)
	t3 := newTreeNode(3)
	t4 := newTreeNode(4)
	t5 := newTreeNode(5)
	t6 := newTreeNode(6)
	t7 := newTreeNode(7)
	t8 := newTreeNode(8)
	t9 := newTreeNode(9)

	t1.setLeft(t2)
	t1.setRight(t3)
	t2.setLeft(t4)
	t2.setRight(t5)
	t3.setLeft(t6)
	t3.setRight(t7)
	t4.setLeft(t8)
	t4.setRight(t9)

	vals := []int{1, 2, 4, 8, 9, 5, 3, 6, 7}
	for i, v := range preorderTraversal(t1) {
		if vals[i] != v {
			t.Errorf("%d %d %d", i, v, vals[i])
			return
		}
	}
}

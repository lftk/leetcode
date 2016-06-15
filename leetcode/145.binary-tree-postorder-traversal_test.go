package leetcode

import "testing"

func Test_postorderTraversal(t *testing.T) {
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

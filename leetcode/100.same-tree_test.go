package leetcode

import "testing"

func Test_isSameTree(t *testing.T) {
	p1 := &TreeNode{Val: 1}
	p2 := &TreeNode{Val: 2}
	p3 := &TreeNode{Val: 3}
	p1.Right = p2
	p2.Left = p3

	q1 := &TreeNode{Val: 1}
	q2 := &TreeNode{Val: 2}
	q3 := &TreeNode{Val: 3}
	q1.Right = q2
	q2.Left = q3

	b := isSameTree(p1, q1)
	if b == false {
		t.Fatal(p1, q1)
	}
}

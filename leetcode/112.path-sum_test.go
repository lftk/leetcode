package leetcode

import "testing"

func Test_hasPathSum(t *testing.T) {
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}
	n8 := &TreeNode{Val: 8}
	n11 := &TreeNode{Val: 11}
	n13 := &TreeNode{Val: 13}
	n3 := &TreeNode{Val: 3}
	n7 := &TreeNode{Val: 7}
	n2 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1}
	n5.Left = n4
	n5.Right = n8
	n4.Left = n11
	n8.Left = n13
	n8.Right = n3
	n11.Left = n7
	n11.Right = n2
	n3.Right = n1
	if !hasPathSum(n5, 22) {
		t.Fatal("no find 5->4->11->2")
	}
}

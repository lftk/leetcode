package leetcode

import "testing"

func Test_invertTree(t *testing.T) {
	n4 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2}
	n7 := &TreeNode{Val: 7}
	n1 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 3}
	n6 := &TreeNode{Val: 6}
	n9 := &TreeNode{Val: 9}
	n4.Left = n2
	n4.Right = n7
	n2.Left = n1
	n2.Right = n3
	n7.Left = n6
	n7.Right = n9
	n := invertTree(n4)
	var left, right int
	for l := n; l != nil; l = l.Left {
		left = left*10 + l.Val
	}
	if left != 479 {
		t.Fatal(n)
	}
	for r := n; r != nil; r = r.Right {
		right = right*10 + r.Val
	}
	if right != 421 {
		t.Fatal(n)
	}
}

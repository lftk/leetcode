package leetcode

import "testing"

func Test_levelOrder(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}
	n3.Left = n9
	n3.Right = n20
	n20.Left = n15
	n20.Right = n7
	vals := levelOrder(n3)
	if len(vals) != 3 {
		t.Fatal("len(vals) != 3")
	}
	if vals[0][0] != 3 {
		t.Fatal(vals)
	}
	if vals[1][0] != 9 || vals[1][1] != 20 {
		t.Fatal(vals)
	}
	if vals[2][0] != 15 || vals[2][1] != 7 {
		t.Fatal(vals)
	}
}

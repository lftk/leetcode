package leetcode

import "testing"

func Test_pathSum(t *testing.T) {
	n1 := &TreeNode{Val: 5}
	n2 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 8}
	n4 := &TreeNode{Val: 11}
	n5 := &TreeNode{Val: 13}
	n6 := &TreeNode{Val: 4}
	n7 := &TreeNode{Val: 7}
	n8 := &TreeNode{Val: 2}
	n9 := &TreeNode{Val: 5}
	n10 := &TreeNode{Val: 1}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n3.Left = n5
	n3.Right = n6
	n4.Left = n7
	n4.Right = n8
	n6.Left = n9
	n6.Right = n10

	// 5, 4, 11, 2
	// 5, 8, 4, 5
	ret := pathSum(n1, 22)
	if len(ret) != 2 {
		t.Fatal(ret)
	}
	if n := ret[0]; len(n) != 4 || n[0] != 5 || n[3] != 2 {
		t.Fatal(n)
	}
	if n := ret[1]; len(n) != 4 || n[0] != 5 || n[3] != 5 {
		t.Fatal(n)
	}
}

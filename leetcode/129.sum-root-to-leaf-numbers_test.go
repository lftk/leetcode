package leetcode

import "testing"

func Test_sumNumbers(t *testing.T) {
	n := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	sum := sumNumbers(n)
	if sum != 25 {
		t.Fatalf("%d != 25", sum)
	}
}

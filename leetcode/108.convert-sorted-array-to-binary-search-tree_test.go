package leetcode

import "testing"

func Test_sortedArrayToBST(t *testing.T) {
	root := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6})
	if root.Val != 4 {
		t.Fatal(root)
	}
	if root.Left.Val != 2 {
		t.Fatal(root)
	}
}

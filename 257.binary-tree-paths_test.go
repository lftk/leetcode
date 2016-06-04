package leetcode

import (
	"strconv"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	var ss []string
	val := strconv.Itoa(root.Val) + "->"
	if root.Left != nil {
		for _, s := range binaryTreePaths(root.Left) {
			ss = append(ss, val+s)
		}
	}
	if root.Right != nil {
		for _, s := range binaryTreePaths(root.Right) {
			ss = append(ss, val+s)
		}
	}
	return ss
}

func TestBinaryTreePaths(t *testing.T) {
	n := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	s := binaryTreePaths(n)
	if len(s) != 2 {
		t.Fatal("len(s) != 2")
	}
	if s[0] != "1->2" || s[1] != "1->3" {
		t.Fatal(s)
	}
}

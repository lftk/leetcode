package leetcode

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	for _, n := range pathNumber(root) {
		v, _ := strconv.Atoi(n)
		sum += v
	}
	return sum
}

func pathNumber(root *TreeNode) (nums []string) {
	var number = map[int]string{
		0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
	}

	val := number[root.Val]
	if root.Left == nil && root.Right == nil {
		return []string{val}
	}
	if root.Left != nil {
		for _, n := range pathNumber(root.Left) {
			nums = append(nums, val+n)
		}
	}
	if root.Right != nil {
		for _, n := range pathNumber(root.Right) {
			nums = append(nums, val+n)
		}
	}
	return
}

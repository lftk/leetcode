// https://leetcode.com/problems/unique-paths/

package leetcode

func uniquePaths(m int, n int) int {
	nums := make([]int, n)
	nums[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			nums[j] += nums[j-1]
		}
	}
	return nums[n-1]
}

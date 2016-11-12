// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/

package leetcode

func minMoves(nums []int) int {
	sum, min := 0, (1<<31 - 1)
	for _, num := range nums {
		if min > num {
			min = num
		}
		sum += num
	}
	return sum - len(nums)*min
}

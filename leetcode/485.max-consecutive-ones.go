// https://leetcode.com/problems/max-consecutive-ones

package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	var k, max int
	for _, n := range nums {
		if k += n; n == 0 {
			if k > max {
				max = k
			}
			k = 0
		}
	}
	if k > max {
		max = k
	}
	return max
}

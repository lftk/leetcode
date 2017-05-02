// https://leetcode.com/problems/array-partition-i

package leetcode

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	var sum int
	sort.Ints(nums)
	for i, n := range nums {
		if i%2 == 0 {
			sum += n
		}
	}
	return sum
}

// https://leetcode.com/problems/missing-number/

package leetcode

func missingNumber(nums []int) int {
	/*
		n := len(nums)
		s1, s2 := 0, (1+n)*n/2
		for _, n := range nums {
			s1 += n
		}
		return s2 - s1
	*/
	sum := len(nums)
	for i, n := range nums {
		sum += i - n
	}
	return sum
}

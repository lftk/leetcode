// https://leetcode.com/problems/single-number-iii/

package leetcode

func singleNumber3(nums []int) []int {
	if nums == nil {
		return nil
	}
	var n1, n2, bit int
	for _, n := range nums {
		bit ^= n
	}
	bit = bit & ^(bit - 1)
	for _, n := range nums {
		if n&bit == bit {
			n1 ^= n
		} else {
			n2 ^= n
		}
	}
	return []int{n1, n2}
}

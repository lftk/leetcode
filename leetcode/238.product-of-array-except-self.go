// https://leetcode.com/problems/product-of-array-except-self/

package leetcode

func productExceptSelf(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return nil
	}
	tmp1, tmp2 := 1, 1
	vals := make([]int, size)
	for i := 0; i < size; i++ {
		vals[i] = 1
	}
	for i, j := 0, size-1; i < size; i, j = i+1, j-1 {
		vals[i] *= tmp1
		tmp1 *= nums[i]
		vals[j] *= tmp2
		tmp2 *= nums[j]
	}
	return vals
}

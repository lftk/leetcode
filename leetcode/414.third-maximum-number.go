// https://leetcode.com/problems/third-maximum-number/

package leetcode

func thirdMax(nums []int) int {
	var count int
	min := -1<<31 - 1
	vals := []int{min, min, min}
	for _, num := range nums {
		if num == vals[0] || num == vals[1] {
			continue
		}
		if num > vals[0] {
			vals[2] = vals[1]
			vals[1] = vals[0]
			vals[0] = num
			count++
		} else if num > vals[1] {
			vals[2] = vals[1]
			vals[1] = num
			count++
		} else if num >= vals[2] {
			vals[2] = num
			count++
		}
	}
	if count < 3 {
		return vals[0]
	}
	return vals[2]
}

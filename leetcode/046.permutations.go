// https://leetcode.com/problems/permutations/

package leetcode

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var ret [][]int
	for i, n := range nums {
		sub := make([]int, len(nums)-1)
		copy(sub, nums[:i])
		copy(sub[i:], nums[i+1:])
		for _, r := range permute(sub) {
			ret = append(ret, append([]int{n}, r...))
		}
	}
	return ret
}

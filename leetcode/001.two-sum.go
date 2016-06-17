// https://leetcode.com/problems/two-sum/

package leetcode

func twoSum(nums []int, target int) []int {
	is := make(map[int]int)
	for i, n := range nums {
		is[n] = i
	}
	for i, n := range nums {
		if j, ok := is[target-n]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}

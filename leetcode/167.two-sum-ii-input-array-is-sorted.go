// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

package leetcode

func twoSum2(numbers []int, target int) []int {
	for i, n := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			sum := n + numbers[j]
			if sum == target {
				return []int{i + 1, j + 1}
			}
			if sum > target {
				break
			}
		}
	}
	return nil
}

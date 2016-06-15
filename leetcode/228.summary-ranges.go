package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	var ret []string
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for j < len(nums) {
			if nums[j]-nums[i] != j-i {
				break
			}
			j++
		}
		if j--; j == i {
			ret = append(ret, fmt.Sprintf("%d", nums[i]))
		} else {
			ret = append(ret, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		}
		i = j
	}
	return ret
}

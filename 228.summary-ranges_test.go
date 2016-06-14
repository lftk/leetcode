package leetcode

import (
	"fmt"
	"testing"
)

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

func TestSummaryRanges(t *testing.T) {
	s := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	if len(s) != 3 {
		t.Fatal(s)
	}
	if s[0] != "0->2" || s[1] != "4->5" || s[2] != "7" {
		t.Fatal(s)
	}
}

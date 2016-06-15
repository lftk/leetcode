package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	c := Candidates(candidates)
	sort.Sort(c)
	return orderedCombinationSum2(c, target)
}

type Candidates []int

func (c Candidates) Len() int {
	return len(c)
}

func (c Candidates) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Candidates) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func orderedCombinationSum2(c Candidates, target int) [][]int {
	var ret [][]int
	for i := 0; i < c.Len(); i++ {
		n := c[i]
		if n > target {
			break
		} else if n == target {
			ret = append(ret, []int{n})
		} else if c.Len() != 1 {
			if s := orderedCombinationSum2(c[i+1:], target-n); s != nil {
				for j := 0; j < len(s); j++ {
					s[j] = append([]int{n}, s[j]...)
					ret = append(ret, s[j])
				}
			}
		}

		for ; i+1 < c.Len(); i++ {
			if n != c[i+1] {
				break
			}
		}
	}
	return ret
}

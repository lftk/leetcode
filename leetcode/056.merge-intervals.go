package leetcode

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	src := Intervals(intervals)
	sort.Sort(src)

	dst := []Interval{src[0]}

	for i := 0; i < src.Len(); i++ {
		if dst[len(dst)-1].End < src[i].Start {
			dst = append(dst, src[i])
		} else {
			dst[len(dst)-1].End = max(dst[len(dst)-1].End, src[i].End)
		}
	}

	return dst

}

// func max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }

type Intervals []Interval

func (is Intervals) Len() int {
	return len(is)
}

func (is Intervals) Less(i, j int) bool {
	return is[i].Start < is[j].Start
}

func (is Intervals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

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
	if src.Len() == 1 {
		return dst
	}

	i, j := 0, 1
	for {
		for dst[i].End < src[j].Start {
			i++
			dst = append(dst, src[j])
			if j++; j == src.Len() {
				return dst
			}
		}
		for dst[i].End >= src[j].Start {
			if dst[i].End < src[j].End {
				dst[i].End = src[j].End
			}
			if j++; j == src.Len() {
				return dst
			}
		}
	}
}

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

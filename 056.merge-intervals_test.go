package leetcode

import (
	"sort"
	"testing"
)

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

func TestMerge(t *testing.T) {
	i1 := Interval{Start: 1, End: 3}
	i2 := Interval{Start: 2, End: 6}
	i3 := Interval{Start: 8, End: 10}
	i4 := Interval{Start: 15, End: 18}
	ret := merge([]Interval{i1, i2, i3, i4})
	if len(ret) != 3 {
		t.Fatal("len(ret) != 3")
	}
	if ret[0].Start != 1 || ret[0].End != 6 {
		t.Fatal(ret)
	}
}

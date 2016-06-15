package leetcode

import "testing"

func Test_merge(t *testing.T) {
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

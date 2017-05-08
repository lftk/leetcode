package leetcode

import (
	"testing"
)

func Test_constructRectangle(t *testing.T) {
	r1 := constructRectangle(180)
	if r1[0] != 15 || r1[1] != 12 {
		t.Error(r1)
		return
	}
}

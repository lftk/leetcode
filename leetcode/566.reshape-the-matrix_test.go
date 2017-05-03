package leetcode

import (
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	m1 := matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4)
	if len(m1) != 1 || len(m1[0]) != 4 {
		t.Error(m1)
		return
	}
	if m1[0][0] != 1 || m1[0][1] != 2 || m1[0][2] != 3 || m1[0][3] != 4 {
		t.Error(m1)
		return
	}
}

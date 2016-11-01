package leetcode

import (
	"testing"
)

func Test_compareVersion(t *testing.T) {
	n1 := compareVersion("1", "1.2")
	if n1 != -1 {
		t.Error(n1, -1)
	}

	n2 := compareVersion("1.0.1", "1")
	if n2 != 1 {
		t.Error(n2, 1)
	}

	n3 := compareVersion("1.1", "1.2")
	if n3 != -1 {
		t.Error(n3, -1)
	}
}

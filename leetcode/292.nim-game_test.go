package leetcode

import (
	"testing"
)

func Test_canWinNim(t *testing.T) {
	b1 := canWinNim(100)
	if b1 {
		t.Errorf("%d should win nim game", b1)
	}
}

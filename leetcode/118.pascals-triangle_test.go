package leetcode

import "testing"

func Test_generate(t *testing.T) {
	val := generate(5)
	if len(val) != 5 {
		t.Fatalf("len(val) %d != 5", len(val))
	}
	if v := val[4]; v[2] != 6 {
		t.Fatal(val)
	}
}

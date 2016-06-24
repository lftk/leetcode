package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	l1 := []string{
		"[{([])}]",
		"{[]}",
		"({[()]})",
	}
	for _, s := range l1 {
		if !isValid(s) {
			t.Fatal(s)
		}
	}
	l2 := []string{
		"[{({[])}]",
		"{[]]}[",
		"({[([})]})",
		"[",
	}
	for _, s := range l2 {
		if isValid(s) {
			t.Fatal(s)
		}
	}
}

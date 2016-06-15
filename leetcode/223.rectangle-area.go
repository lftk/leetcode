package leetcode

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	s := (C-A)*(D-B) + (G-E)*(H-F)
	if E >= C || A >= G || F >= D || H <= B {
		return s
	}
	x := (min(C, G) - max(A, E))
	y := (min(D, H) - max(B, F))
	return s - x*y
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

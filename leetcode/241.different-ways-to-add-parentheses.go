package leetcode

import "strconv"

func diffWaysToCompute(input string) []int {
	nums, ops := parseString(input)
	exps := createExpNodes(nums, ops)
	return calcExpNodes(exps)
}

type expNode struct {
	val         int
	op          rune
	left, right []*expNode
}

func createExpNodes(nums []int, ops []rune) (exps []*expNode) {
	if len(ops) == 0 {
		exps = append(exps, &expNode{val: nums[0]})
		return
	}

	for i, op := range ops {
		exp := &expNode{
			op:    op,
			left:  createExpNodes(nums[:i+1], ops[:i]),
			right: createExpNodes(nums[i+1:], ops[i+1:]),
		}
		exps = append(exps, exp)
	}
	return
}

func calcExpNodes(exps []*expNode) (nums []int) {
	for _, exp := range exps {
		nums = append(nums, calcExpNode(exp)...)
	}
	return
}

func calcExpNode(exp *expNode) (nums []int) {
	if exp.op == rune(0) {
		nums = append(nums, exp.val)
		return
	}

	left := calcExpNodes(exp.left)
	right := calcExpNodes(exp.right)

	for _, l := range left {
		for _, r := range right {
			nums = append(nums, opNumber(l, r, exp.op))
		}
	}
	return
}

func parseString(input string) (nums []int, ops []rune) {
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+', '-', '*':
			num, _ := strconv.Atoi(input[:i])
			op := rune(input[i])
			nums = append(nums, num)
			ops = append(ops, op)
			input = input[i+1:]
			i = 0
		}
	}
	num, _ := strconv.Atoi(input)
	nums = append(nums, num)
	return
}

func opNumber(n1, n2 int, op rune) (n int) {
	switch op {
	case '+':
		n = n1 + n2
	case '-':
		n = n1 - n2
	case '*':
		n = n1 * n2
	}
	return
}

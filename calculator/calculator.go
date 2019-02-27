package calculator

import (
	"strconv"
	"strings"
)

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Eval(expr string) int {
	var ops []string
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if strings.Index(higher, op) < 0 {
				// 목록에 없는 연산자
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				return
			}
			b, a := pop(), pop()
			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}
	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			reduce("+-*/")
			ops = append(ops, token)
		case ")":
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}

	reduce("+-*/")
	return nums[0]
}

package main

import "strconv"

//https://leetcode.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stack := make([]int64, 0, len(tokens))
	for _, c := range tokens {
		num, err := strconv.ParseInt(c, 10, 64)
		if err == nil {
			stack = append(stack, num)
			continue
		}
		op1 := stack[len(stack)-2]
		op2 := stack[len(stack)-1]
		result := int64(0)
		switch c {
		case "+":
			result = op1 + op2
		case "-":
			result = op1 - op2
		case "*":
			result = op1 * op2
		case "/":
			result = op1 / op2
		}
		stack[len(stack)-2] = result
		stack = stack[:len(stack)-1]
	}
	return int(stack[0])
}

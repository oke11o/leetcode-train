package _6xx

import "strconv"

/**
https://leetcode.com/problems/baseball-game/
682. Baseball Game
Easy
*/
func calPoints(ops []string) int {
	stack := []int{}
	for _, v := range ops {
		switch v {
		case "+":
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			newtop := top + stack[len(stack)-1]
			stack = append(stack, top)
			stack = append(stack, newtop)
		case "C":
			stack = stack[:len(stack)-1] // pop
		case "D":
			stack = append(stack, 2*stack[len(stack)-1])
		default:
			d, _ := strconv.Atoi(v)
			stack = append(stack, d)
		}
	}
	var ans int
	for _, v := range stack {
		ans += v
	}

	return ans
}

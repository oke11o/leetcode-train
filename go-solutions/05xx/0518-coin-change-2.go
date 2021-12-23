package _5xx

// 0518. Coin Change 2
//  Medium
// #Amazon
// #array, #dynamic_programming, #need_explaining
// https://leetcode.com/problems/coin-change-2/
// https://www.educative.io/blog/crack-amazon-coding-interview-questions
// https://www.educative.io/m/coin-changing-problem
func change(amount int, coins []int) int {
	solution := make([]int, amount+1)
	solution[0] = 1
	for _, den := range coins {
		for i := den; i < amount+1; i++ {
			solution[i] += solution[i-den]
		}
	}
	return solution[len(solution)-1]
}

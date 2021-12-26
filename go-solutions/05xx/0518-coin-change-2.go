package _5xx

/**
 * 0518. Coin Change 2
 * Medium
 * #Amazon
 * #array, #dynamic_programming, #need_explaining
 * https://leetcode.com/problems/coin-change-2/
 * https://www.educative.io/blog/crack-amazon-coding-interview-questions
 * https://www.educative.io/m/coin-changing-problem
 */
func change(amount int, coins []int) int {
	solution := make([]int, amount+1) // solution[i] - сколько монет можно положить, чтобы получуть сумму i
	solution[0] = 1                   // чтобы потом на первом проходе во все ячейки поставить 1. Так как, если у нас есть монета 1, то из нее можно любое кол-во сумм составить
	for _, den := range coins {
		for sum := den; sum < amount+1; sum++ { // Начинаем со позиции den, так как ниже будет sum-den, которая не может быть менее 0
			solution[sum] += solution[sum-den]
		}
	}
	return solution[amount]
}

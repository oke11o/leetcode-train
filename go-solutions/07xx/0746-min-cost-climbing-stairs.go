package _7xx

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt  = 1<<(intSize-1) - 1
	MinInt  = -1 << (intSize - 1)
)

// 0746. Min Cost Climbing Stairs
func minCostClimbingStairs(cost []int) int {
	var min = func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

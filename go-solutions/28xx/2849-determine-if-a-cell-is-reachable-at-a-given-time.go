package _8xx

/*
2849. Determine if a Cell Is Reachable at a Given Time
https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/?envType=daily-question&envId=2023-11-08
*/
func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	x := sx - fx
	if x < 0 {
		x = -x
	}
	y := sy - fy
	if y < 0 {
		y = -y
	}
	m := max(x, y)
	if m == 0 {
		return t != 1
	}

	return t >= m
}

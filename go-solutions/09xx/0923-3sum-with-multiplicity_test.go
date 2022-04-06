package _9xx

import "testing"

/**
https://leetcode.com/problems/3sum-with-multiplicity/
923. 3Sum With Multiplicity
Medium
#two_pointers
*/

/**
Approach 3.
По сути, надо разобраться с теми сложными условиями, которые надо выполнить, чтобы решить проблему повторяющихся значений
*/
func threeSumMulti(arr []int, target int) int {
	mod := 1000000007

	// Initializing as long saves us the trouble of
	// managing count[x] * count[y] * count[z] overflowing later.
	count := make([]int, 101)
	uniq := 0
	for _, v := range arr {
		count[v]++
		if count[v] == 1 {
			uniq++
		}
	}

	keys := make([]int, 0, uniq)
	for i := 0; i <= 100; i++ {
		if count[i] > 0 {
			keys = append(keys, i)
		}
	}

	ans := 0

	// Now, let's do a 3sum on "keys", for i <= j <= k.
	// We will use count to add the correct contribution to ans.
	for i := 0; i < uniq; i++ {
		x := keys[i]
		T := target - x
		j := i
		k := uniq - 1
		for j <= k {
			y := keys[j]
			z := keys[k]
			if y+z < T {
				j++
			} else if y+z > T {
				k--
			} else { // # x+y+z == T, now calc the size of the contribution
				if i < j && j < k {
					ans += count[x] * count[y] * count[z]
				} else if i == j && j < k {
					ans += count[x] * (count[x] - 1) / 2 * count[z]
				} else if i < j && j == k {
					ans += count[x] * count[y] * (count[y] - 1) / 2
				} else { // i == j == k
					ans += count[x] * (count[x] - 1) * (count[x] - 2) / 6
				}
				ans %= mod
				j++
				k--
			}
		}
	}

	return ans
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_threeSumMulti(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumMulti(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("threeSumMulti() = %v, want %v", got, tt.want)
			}
		})
	}
}

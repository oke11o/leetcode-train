package _3xx

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

/**
https://leetcode.com/problems/top-k-frequent-elements/
347. Top K Frequent Elements
Medium
*/
func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num] += 1
	}

	uniqNums := make([]int, 0, len(counter))
	for num := range counter {
		uniqNums = append(uniqNums, num)
	}

	sort.Slice(uniqNums, func(i, j int) bool {
		return counter[uniqNums[i]] > counter[uniqNums[j]]
	})

	return uniqNums[:k]
	numFreq := make(map[int]int)
	for _, n := range nums {
		numFreq[n]++
	}

	freqs := make([]int, 0, len(numFreq))
	freqNum := make(map[int][]int, len(numFreq))
	for n, v := range numFreq {
		freqNum[v] = append(freqNum[v], n)
		if len(freqNum[v]) == 1 {
			freqs = append(freqs, v)
		}
	}

	sort.Ints(freqs)
	result := make([]int, 0, k)
	for i := len(freqs) - 1; i >= 0 && k > 0; i-- {
		for _, v := range freqNum[freqs[i]] {
			result = append(result, v)
			k--
			if k == 0 {
				break
			}
		}
	}

	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
				k:    10,
			},
			want: []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.args.nums, tt.args.k)
			require.ElementsMatch(t, got, tt.want)
		})
	}
}

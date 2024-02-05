package _3xx

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
https://leetcode.com/problems/check-if-n-and-its-double-exist/solutions/4147694/a-different-c-solution-o-1-space-two-pointers-sorting/
1346. Check If N and Its Double Exist
Easy
#two_pointer, #array, #sorting, #hash_table, #binary_search
*/
func checkIfExist(arr []int) bool {
	sort.Ints(arr)

	i := 0
	j := 0
	n := len(arr)
	for i < n && j < n {
		for j < n && arr[j] < 2*arr[i] {
			j++
		}
		if !(j < n) {
			break
		}
		if arr[j] == 2*arr[i] && i != j {
			return true
		}
		i++
	}

	return false
}

func Test_checkIfExist(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "",
			arr:  []int{10, 2, 5, 3},
			want: true,
		},
		{
			name: "",
			arr:  []int{-10, 2, 5, 3},
			want: false,
		},
		{
			name: "",
			arr:  []int{-11, -13, -10, 2, -5, 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkIfExist(tt.arr)
			require.Equal(t, tt.want, got)
		})
	}
}

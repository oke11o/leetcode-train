package _7xx

import (
	"sort"
	"testing"
)

/**
https://leetcode.com/problems/maximum-units-on-a-truck/
1710. Maximum Units on a Truck
Easy
- numberOfBoxesi is the number of boxes of type i.
- numberOfUnitsPerBoxi is the number of units in each box of the type i.

Similars
- https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/
- https://leetcode.com/problems/partition-array-into-disjoint-intervals/
- https://leetcode.com/problems/minimum-incompatibility/
*/
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	result := 0
	i := 0
	for truckSize > 0 && i < len(boxTypes) {
		boxCount := truckSize
		if boxCount > boxTypes[i][0] {
			boxCount = boxTypes[i][0]
		}
		result += boxCount * boxTypes[i][1]
		truckSize -= boxCount
		i++
	}
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_maximumUnits(t *testing.T) {
	type args struct {
		boxTypes  [][]int
		truckSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				boxTypes:  [][]int{{1, 3}, {2, 2}, {3, 1}},
				truckSize: 4,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				boxTypes:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
				truckSize: 10,
			},
			want: 91,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUnits(tt.args.boxTypes, tt.args.truckSize); got != tt.want {
				t.Errorf("maximumUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}

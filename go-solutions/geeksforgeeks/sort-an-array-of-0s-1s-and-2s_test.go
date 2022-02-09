package geeksforgeeks

import (
	"reflect"
	"testing"
)

/**
Sort an array of 0s, 1s and 2s
Medium
https://www.geeksforgeeks.org/sort-an-array-of-0s-1s-and-2s/
[#sliding_window]
*/
func sort012(flag []int) {
	zeroP := 0
	twoP := len(flag) - 1
	for zeroP < twoP && flag[zeroP] == 0 {
		zeroP++
	}
	if zeroP == twoP {
		return
	}
	for twoP >= 0 && flag[twoP] == 2 {
		twoP--
	}
	if twoP == 0 {
		return
	}
	pointer := zeroP
	for pointer <= twoP {
		if flag[pointer] == 0 {
			flag[pointer], flag[zeroP] = flag[zeroP], flag[pointer]
			zeroP++
		} else if flag[pointer] == 2 {
			flag[pointer], flag[twoP] = flag[twoP], flag[pointer]
			twoP--
		} else {
			pointer++
		}
	}
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_sort012(t *testing.T) {
	tests := []struct {
		name string
		flag []int
		want []int
	}{
		{
			name: "",
			flag: []int{0, 1, 2, 0, 1, 2},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "",
			flag: []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1},
			want: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2},
		},
		{
			name: "",
			flag: []int{0, 0, 0, 0, 0, 0},
			want: []int{0, 0, 0, 0, 0, 0},
		},
		{
			name: "",
			flag: []int{1, 1, 1, 1, 1, 1},
			want: []int{1, 1, 1, 1, 1, 1},
		},
		{
			name: "",
			flag: []int{2, 2, 2, 2, 2, 2},
			want: []int{2, 2, 2, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort012(tt.flag)
			if !reflect.DeepEqual(tt.flag, tt.want) {
				t.Errorf("twoSum() = %v, want %v", tt.flag, tt.want)
			}
		})
	}
}

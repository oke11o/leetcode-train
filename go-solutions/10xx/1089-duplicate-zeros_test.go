package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
*
https://leetcode.com/problems/duplicate-zeros/
1089. Duplicate Zeros
*/
func duplicateZeros(arr []int) {
	possibleDups := 0
	length_ := len(arr) - 1

	// Find the number of zeros to be duplicated
	// Stopping when left points beyond the last element in the original array
	// which would be part of the modified array
	for left := 0; left <= length_-possibleDups; left++ {

		// Count the zeros
		if arr[left] == 0 {

			// Edge case: This zero can't be duplicated. We have no more space,
			// as left is pointing to the last element which could be included
			if left == length_-possibleDups {
				// For this zero we just copy it without duplication.
				arr[length_] = 0
				length_ -= 1
				break
			}
			possibleDups++
		}
	}

	// Start backwards from the last element which would be part of new array.
	last := length_ - possibleDups

	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+possibleDups] = 0
			possibleDups--
			arr[i+possibleDups] = 0
		} else {
			arr[i+possibleDups] = arr[i]
		}
	}
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_duplicateZeros(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "",
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "",
			arr:  []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "",
			arr:  []int{8, 4, 5, 0, 0, 0, 0, 7},
			want: []int{8, 4, 5, 0, 0, 0, 0, 0},
		},
		{
			name: "",
			arr:  []int{0, 0, 0, 0, 0},
			want: []int{0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.arr)
			require.Equal(t, tt.want, tt.arr)
		})
	}
}

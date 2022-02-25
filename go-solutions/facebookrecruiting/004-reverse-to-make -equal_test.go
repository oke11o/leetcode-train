package facebookrecruiting

import (
	"fmt"
	"testing"
)

/**
Reverse to Make Equal
Given two arrays A and B of length N, determine if there is a way to make A equal to B by reversing any subarrays from array B any number of times.
Signature
bool areTheyEqual(int[] arr_a, int[] b)
Input
All integers in array are in the range [0, 1,000,000,000].
Output
Return true if B can be made equal to A, return false otherwise.
Example
A = [1, 2, 3, 4]
B = [1, 4, 3, 2]
output = true
After reversing the subarray of B from indices 1 to 3, array B will equal array A.
*/
func areTheyEqual(a []int, b []int) bool {
	hash := make(map[int]int, len(a))
	for i := 0; i < len(a); i++ {
		hash[a[i]]++
		hash[b[i]]++
	}
	for _, cnt := range hash {
		if cnt%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct {
		input1 []int
		input2 []int
		want   bool
	}{
		{
			input1: []int{1, 2, 3, 4},
			input2: []int{1, 4, 3, 2},
			want:   true,
		},
	}
	for i, tt := range tests {
		got := areTheyEqual(tt.input1, tt.input2)
		if got == tt.want {
			fmt.Printf("#%d OK\n", i)
		} else {
			fmt.Printf("#%d invalid got/want `%v` != `%v`\n", i, got, tt.want)
		}
	}
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_areTheyEqual(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				a: []int{1, 2, 3, 4},
				b: []int{1, 4, 3, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areTheyEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("areTheyEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

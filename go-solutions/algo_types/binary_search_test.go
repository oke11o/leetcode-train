package algo_types

import "testing"

/**
Binary search
*/
func binarySearch(arr []int, x int) int {
	left := -1
	right := len(arr)
	for (right - left) > 1 {
		i := (left + right) / 2
		if arr[i] <= x {
			left = i
		} else {
			right = i
		}
	}
	if left == -1 {
		return 0
	}

	return left
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_binarySearch(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{2, 10, 10, 45, 46, 51, 66, 71, 73, 106, 157, 187, 212, 283, 303, 316, 398, 405, 425, 464, 477, 508, 590, 600, 604, 771, 828},
				x:   57,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

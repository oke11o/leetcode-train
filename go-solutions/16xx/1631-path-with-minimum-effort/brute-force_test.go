package _631_path_with_minimum_effort

import "testing"

func Test_minimumEffortPath(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name        string
		args        args
		explanation string
		want        int
	}{
		{
			name: "",
			args: args{
				heights: [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			},
			explanation: `The route of {1,3,5,3,5} has a maximum absolute difference of 2 in consecutive cells.
This is better than the route of {1,2,2,2,5}, where the maximum absolute difference is 3.`,
			want: 2,
		},
		{
			name: "",
			args: args{
				heights: [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			},
			explanation: `The route of {1,2,3,4,5} has a maximum absolute difference of 1 in consecutive cells, which is better than route {1,3,5,3,5}.`,
			want:        1,
		},
		{
			name: "",
			args: args{
				heights: [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}},
			},
			explanation: `This route does not require any effort`,
			want:        0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPath(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

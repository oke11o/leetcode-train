package main

import "testing"

func Test_problemD(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		arr  []int
		want string
	}{
		{
			name: "",
			arr:  []int{1, 2, 3, 4, 5},
			want: "YES",
		},
		{
			name: "",
			arr:  []int{1, 2, 3, 1},
			want: "NO",
		},
		{
			name: "",
			arr:  []int{2, 3, 4, 8, 5, 5, 5, 5},
			want: "YES",
		},
		{
			name: "",
			arr:  []int{1, 1, 3, 2, 2},
			want: "YES",
		},
		{
			name: "",
			arr:  []int{1, 1, 2, 3, 2},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problemD(tt.arr); got != tt.want {
				t.Errorf("problemD() = %v, want %v", got, tt.want)
			}
		})
	}
}

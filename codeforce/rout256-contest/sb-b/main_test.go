package main

import "testing"

func Test_problemB(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "",
			arr:  []int{2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3},
			want: 22,
		},
		{
			name: "",
			arr:  []int{2, 3, 2, 3, 2, 2, 3, 2, 3, 2, 2, 3},
			want: 22,
		},
		{
			name: "",
			arr:  []int{10000},
			want: 10000,
		},
		{
			name: "",
			arr:  []int{1, 2, 3, 1, 2, 3, 1, 2, 3},
			want: 12,
		},
		{
			name: "",
			arr:  []int{10000, 10000, 10000, 10000, 10000, 10000},
			want: 40000,
		},
		{
			name: "",
			arr:  []int{300, 100, 200, 300, 200, 300},
			want: 1100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problemB(tt.arr); got != tt.want {
				t.Errorf("problemB() = %v, want %v", got, tt.want)
			}
		})
	}
}

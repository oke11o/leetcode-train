package main

import (
	"reflect"
	"testing"
)

func Test_problemC(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want [][2]int
	}{
		{
			name: "",
			arr:  []int{2, 1, 3, 1, 1, 4},
			want: [][2]int{{1, 2}, {3, 6}, {4, 5}},
		},
		{
			name: "",
			arr:  []int{5, 5},
			want: [][2]int{{1, 2}},
		},
		{
			name: "",
			arr:  []int{1, 4, 2, 5, 4, 2, 6, 3},
			want: [][2]int{{1, 3}, {2, 5}, {4, 7}, {6, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problemC(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("problemC() = %v, want %v", got, tt.want)
			}
		})
	}
}

package _5xx

import "testing"

type Node struct {
	Val      int
	Children []*Node
}

/**
https://leetcode.com/problems/diameter-of-n-ary-tree/submissions/
1522. Diameter of N-Ary Tree
Medium

*/
var diameterResult int

func diameter(root *Node) int {
	diameterResult = 0
	longestPath(root)
	return diameterResult
}

func longestPath(node *Node) int {
	if node == nil {
		return 0
	}
	if len(node.Children) == 0 {
		return 1
	}

	paths := make([]int, len(node.Children))
	for i := 0; i < len(paths); i++ {
		paths[i] = longestPath(node.Children[i])
	}
	diameter, maxPath := findDiameterAndMaxPath(paths)
	if diameterResult < diameter {
		diameterResult = diameter
	}
	return maxPath
}

func findDiameterAndMaxPath(list []int) (diameter int, maxPath int) {
	if len(list) == 1 {
		return list[0], list[0] + 1
	}
	first, second := twoMax(list)
	return first + second, first + 1
}

func twoMax(list []int) (first int, second int) {
	for _, v := range list {
		if v > first {
			second = first
			first = v
		} else if v > second {
			second = v
		}
	}
	return
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_twoMax(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name       string
		args       args
		wantFirst  int
		wantSecond int
	}{
		{
			name: "",
			args: args{
				list: []int{2, 1, 1},
			},
			wantFirst:  2,
			wantSecond: 1,
		},
		{
			name: "",
			args: args{
				list: []int{2, 1, 1, 1},
			},
			wantFirst:  2,
			wantSecond: 1,
		},
		{
			name: "",
			args: args{
				list: []int{2, 1, 2, 1},
			},
			wantFirst:  2,
			wantSecond: 2,
		},
		{
			name: "",
			args: args{
				list: []int{2, 0, 0, 0},
			},
			wantFirst:  2,
			wantSecond: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotSecond := twoMax(tt.args.list)
			if gotFirst != tt.wantFirst {
				t.Errorf("twoMax() gotFirst = %v, want %v", gotFirst, tt.wantFirst)
			}
			if gotSecond != tt.wantSecond {
				t.Errorf("twoMax() gotSecond = %v, want %v", gotSecond, tt.wantSecond)
			}
		})
	}
}
func Test_diameter(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{
							Val: 2,
							Children: []*Node{
								{
									Val:      3,
									Children: []*Node{{Val: 5}},
								},
								{
									Val:      4,
									Children: []*Node{{Val: 6}},
								},
							},
						},
					},
				},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				root: &Node{
					Val: 44,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diameter(tt.args.root)
			if got != tt.want {
				t.Errorf("diameter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

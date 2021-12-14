package _1xx

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func Test_connect(t *testing.T) {
	const Null = math.MinInt32
	var buildSlice = func(root *Node) []int {
		result := make([]int, 0)
		queue := make([]*Node, 0)
		queue = append(queue, root)
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			result = append(result, node.Val)
			if len(queue) == 0 {
				result = append(result, Null)
			}
			if node.Left != nil && node.Right != nil {
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}

		return result
	}
	nodes := []*Node{
		0: nil,
		1: {Val: 1},
		2: {Val: 2},
		3: {Val: 3},
		4: {Val: 4},
		5: {Val: 5},
		6: {Val: 6},
		7: {Val: 7},
	}
	nodes[1].Left = nodes[2]
	nodes[1].Right = nodes[3]
	nodes[2].Left = nodes[4]
	nodes[2].Right = nodes[5]
	nodes[3].Left = nodes[5]
	nodes[3].Right = nodes[7]

	want := []*Node{
		0: nil,
		1: {Val: 1},
		2: {Val: 2},
		3: {Val: 3},
		4: {Val: 4},
		5: {Val: 5},
		6: {Val: 6},
		7: {Val: 7},
	}
	want[1].Left = want[2]
	want[1].Right = want[3]
	want[2].Left = want[4]
	want[2].Right = want[5]
	want[2].Next = want[3]
	want[3].Left = want[5]
	want[3].Right = want[7]
	want[4].Next = want[5]
	want[5].Next = want[6]
	want[6].Next = want[7]

	got := connect(nodes[1])
	slice := buildSlice(want[1])
	require.Equal(t, slice, buildSlice(got))
}

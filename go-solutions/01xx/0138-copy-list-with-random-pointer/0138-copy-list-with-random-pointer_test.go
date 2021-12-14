package _138_copy_list_with_random_pointer

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	const Null = math.MinInt32

	var nodex2list = func(h *Node) [][2]int {
		result := [][2]int{}
		for h != nil {

			v := [2]int{h.Val}
			if h.Random == nil {
				v[1] = Null
			} else {
				v[1] = h.Random.Val
			}
			result = append(result, v)

			h = h.Next
		}

		return result
	}

	tests := []struct {
		name string
		head func() *Node
	}{
		{
			name: "",
			head: func() *Node {
				Node7 := &Node{Val: 7}
				Node13 := &Node{Val: 13}
				Node11 := &Node{Val: 11}
				Node10 := &Node{Val: 10}
				Node1 := &Node{Val: 1}

				Node7.Next = Node13
				Node7.Random = nil
				Node13.Next = Node11
				Node13.Random = Node7
				Node11.Next = Node10
				Node11.Random = Node1
				Node10.Next = Node1
				Node10.Random = Node11
				Node1.Next = nil
				Node1.Random = Node7
				return Node7
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := tt.head()
			want := nodex2list(head)
			node := copyRandomList(head)
			for head != nil {
				head.Val = 0
				head = head.Next
			}
			got := nodex2list(node)
			require.Equal(t, want, got)
		})
	}
}

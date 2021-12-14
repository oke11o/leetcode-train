package _1xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions"
	"github.com/stretchr/testify/require"
	"testing"
)

var sameTreeTests = []struct {
	name string
	p    []int
	q    []int
	want bool
}{
	{
		name: "",
		p:    []int{1, 2, 3},
		q:    []int{1, 2, 3},
		want: true,
	},
	{
		name: "",
		p:    []int{1, 2},
		q:    []int{1, Null, 2},
		want: false,
	},
	{
		name: "",
		p:    []int{1, 2, 1},
		q:    []int{1, 1, 2},
		want: false,
	},
	{
		name: "",
		p:    []int{1, Null, 1, Null, 1, Null, 1, Null, 1, Null, 1, 2},
		q:    []int{1, Null, 1, Null, 1, Null, 1, Null, 1, Null, 1, 2},
		want: true,
	},
}

func Test_isSameTree(t *testing.T) {
	for _, tt := range sameTreeTests {
		t.Run(tt.name, func(t *testing.T) {
			p := CreateTreeNodeFromSlice(tt.p)
			q := CreateTreeNodeFromSlice(tt.q)

			got := isSameTree(p, q)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_isSameTree_Rec(t *testing.T) {
	for _, tt := range sameTreeTests {
		t.Run(tt.name, func(t *testing.T) {
			p := CreateTreeNodeFromSlice(tt.p)
			q := CreateTreeNodeFromSlice(tt.q)

			got := isSameTree_Rec(p, q)
			require.Equal(t, tt.want, got)
		})
	}
}

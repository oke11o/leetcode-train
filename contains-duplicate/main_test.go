package contains_duplicate

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name:"",
			nums:[]int{1,2,3,1},
			want: true,
		},
		{
			name:"",
			nums:[]int{1,2,3,4},
			want: false,
		},
		{
			name:"",
			nums:[]int{1,1,1,3,3,4,3,2,4,2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}

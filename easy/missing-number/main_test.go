package missing_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	name string
	nums []int
	want int
}{
	{
		name: "",
		nums: []int{3, 0, 1},
		want: 2,
	},
	{
		name: "",
		nums: []int{0, 1},
		want: 2,
	},
	{
		name: "",
		nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
		want: 8,
	},
	{
		name: "",
		nums: []int{0},
		want: 1,
	},
}

func Test_missingNumber(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingNumber(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_missingNumber2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingNumber2(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}

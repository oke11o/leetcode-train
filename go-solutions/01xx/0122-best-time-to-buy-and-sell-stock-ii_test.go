package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxProfit_ii(t *testing.T) {
	tests := []struct {
		name        string
		prices      []int
		want        int
		explanation string
	}{
		{
			name:   "",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
			explanation: `Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.`,
		},
		{
			name:   "",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
			explanation: `Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.`,
		},
		{
			name:        "",
			prices:      []int{7, 6, 4, 3, 1},
			want:        0,
			explanation: `There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit_ii(tt.prices)
			require.Equal(t, tt.want, got)
		})
	}
}

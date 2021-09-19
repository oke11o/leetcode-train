package google_kick_start_2021_round_f

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_minDistanceToTrashBin(t *testing.T) {
	tests := []struct {
		name string
		mask string
		want int
	}{
		{
			name: "",
			mask: "111",
			want: 0,
		},
		{
			name: "",
			mask: "100100",
			want: 5,
		},
		{
			name: "",
			mask: "1001000",
			want: 8,
		},
		{
			name: "",
			mask: "100001000",
			want: 12,
		},
		{
			name: "",
			mask: "1001001",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistanceToTrashBin(tt.mask)
			require.Equal(t, tt.want, got)
		})
	}
}

package _2xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals []*Interval
		want      bool
	}{
		{
			name:      "",
			intervals: []*Interval{{0, 30}, {5, 10}, {15, 20}},
			want:      false,
		},
		{
			name:      "",
			intervals: []*Interval{{5, 8}, {9, 15}},
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canAttendMeetings(tt.intervals)
			require.Equal(t, tt.want, got)
		})
	}
}

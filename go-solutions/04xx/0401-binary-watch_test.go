package _4xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	tests := []struct {
		name     string
		turnedOn int
		want     []string
	}{
		{
			name:     "",
			turnedOn: 0,
			want:     []string{"0:00"},
		},
		{
			name:     "",
			turnedOn: 1,
			want:     []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			name:     "",
			turnedOn: 2,
			want:     []string{"10:00", "9:00", "8:32", "8:16", "8:08", "8:04", "8:02", "8:01", "6:00", "5:00", "4:32", "4:16", "4:08", "4:04", "4:02", "4:01", "3:00", "2:32", "2:16", "2:08", "2:04", "2:02", "2:01", "1:32", "1:16", "1:08", "1:04", "1:02", "1:01", "0:48", "0:40", "0:36", "0:34", "0:33", "0:24", "0:20", "0:18", "0:17", "0:12", "0:10", "0:09", "0:06", "0:05", "0:03"},
		},
		{
			name:     "",
			turnedOn: 8,
			want:     []string{"7:31", "7:47", "7:55", "7:59", "11:31", "11:47", "11:55", "11:59"},
		},
		{
			name:     "",
			turnedOn: 9,
			want:     []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readBinaryWatch2(tt.turnedOn)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_watch2string(t *testing.T) {
	tests := []struct {
		name    string
		current int16
		want    string
	}{
		{
			name:    "",
			current: 1,
			want:    "0:01",
		},
		{
			name:    "",
			current: 2,
			want:    "0:02",
		},
		{
			name:    "",
			current: 65,
			want:    "1:01",
		},
		{
			name:    "",
			current: 129,
			want:    "2:01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := watch2string(tt.current)
			require.Equal(t, tt.want, got)
		})
	}
}
func Test_countBits(t *testing.T) {
	tests := []struct {
		name    string
		current int
		want    int
	}{
		{
			name:    "",
			current: 0,
			want:    0,
		},
		{
			name:    "",
			current: 1,
			want:    1,
		},
		{
			name:    "",
			current: 2,
			want:    1,
		},
		{
			name:    "",
			current: 6,
			want:    2,
		},
		{
			name:    "",
			current: 7,
			want:    3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.current)
			require.Equal(t, tt.want, got)
		})
	}

}

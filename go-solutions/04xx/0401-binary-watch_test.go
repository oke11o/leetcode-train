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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readBinaryWatch(tt.turnedOn)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}

package _0xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/powx-n/
50. Pow(x, n)
Medium
*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var res float64 = 1
	for n > 0 {
		if n%2 == 1 {
			res *= x // 4
		}
		x *= x // 16
		n /= 2 // 2
	}
	return res
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				x: 2,
				n: 10,
			},
			want: 1024,
		},
		{
			name: "",
			args: args{
				x: 2.1,
				n: 3,
			},
			want: 9.261000000000001,
		},
		{
			name: "",
			args: args{
				x: 2,
				n: -2,
			},
			want: 0.25,
		},
		{
			name: "",
			args: args{
				x: 0.44528,
				n: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, myPow(tt.args.x, tt.args.n), "myPow(%v, %v)", tt.args.x, tt.args.n)
		})
	}
}

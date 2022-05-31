package _461

import "testing"

func Test_hasAllCodes(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "00110110",
				k: 2,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "0110",
				k: 1,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "0110",
				k: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAllCodes_hashFun(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("hasAllCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

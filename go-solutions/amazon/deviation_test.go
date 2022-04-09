package amazon

import "testing"

func Test_deviations(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "",
			args: args{
				s: "bbacccabab",
			},
			want: 2,
		},
		{
			name: "tc 2",
			args: args{
				s: "abdbcdacbcadbbc",
			},
			want: 3,
		},
		{
			name: "tc 3",
			args: args{
				s: "cccaabbccc",
			},
			want: 4,
		},
		{
			name: "tc 6",
			args: args{
				s: "baccabcbc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deviations(tt.args.s); got != tt.want {
				t.Errorf("deviations() = %v, want %v", got, tt.want)
			}
		})
	}
}

package yandex

import "testing"

func Test_marshal(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want string
	}{
		{
			name: "",
			expr: "",
			want: "",
		},
		{
			name: "",
			expr: "asdf",
			want: "asdf",
		},
		{
			name: "",
			expr: "(ab)[3]",
			want: "ababab",
		},
		{
			name: "",
			expr: "((a)[2]b)[3]",
			want: "aabaabaab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := marshal(tt.expr); got != tt.want {
				t.Errorf("marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

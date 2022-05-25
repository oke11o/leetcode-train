package yandex

import "testing"

func Test_findAnag(t *testing.T) {
	type args struct {
		in   []rune
		need []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				in:   []rune("Hello my new world"),
				need: []rune("ldr"),
			},
			want: 15,
		},
		{
			name: "",
			args: args{
				in:   []rune("Hello my new world"),
				need: []rune("o m"),
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				in:   []rune("Hello my new world"),
				need: []rune("Hel"),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				in:   []rune("Hello my new world"),
				need: []rune("sdfsdf"),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnag(tt.args.in, tt.args.need); got != tt.want {
				t.Errorf("findAnag() = %v, want %v", got, tt.want)
			}
		})
	}
}

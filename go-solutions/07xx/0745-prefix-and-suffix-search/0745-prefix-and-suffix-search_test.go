package _745_prefix_and_suffix_search

import "testing"

func TestWordFilter_F(t *testing.T) {
	type args struct {
		prefix string
		suffix string
	}
	tests := []struct {
		name  string
		words []string
		args  []args
		want  []int
	}{
		{
			name:  "",
			words: []string{"apple"},
			args:  []args{{"a", "e"}},
			want:  []int{0},
		},
		{
			name:  "",
			words: []string{"cabaabaaaa", "ccbcababac", "bacaabccba", "bcbbcbacaa", "abcaccbcaa", "accabaccaa", "cabcbbbcca", "ababccabcb", "caccbbcbab", "bccbacbcba"},
			args: []args{
				{"bccbacbcba", "a"},    //9
				{"ab", "abcaccbcaa"},   //4
				{"a", "aa"},            //5
				{"cabaaba", "abaaaa"},  //0
				{"cacc", "accbbcbab"},  //8
				{"ccbcab", "bac"},      //1
				{"bac", "cba"},         //2
				{"ac", "accabaccaa"},   //5
				{"bcbb", "aa"},         //3
				{"ccbca", "cbcababac"}, //1
			},
			want: []int{9, 4, 5, 0, 8, 1, 2, 5, 3, 1},
		},
		{
			name:  "",
			words: []string{"abc", "abd", "ab"},
			args: []args{
				{"c", "d"}, //-1
				{"a", "c"}, //0
				{"a", "d"}, //1
				{"a", "b"}, //1
			},
			want: []int{-1, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.words)
			for i, arg := range tt.args {
				got := this.F(arg.prefix, arg.suffix)
				if got != tt.want[i] {
					t.Errorf("F() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}

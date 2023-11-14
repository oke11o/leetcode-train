package _9xx

import "testing"

/*
1930. Unique Length-3 Palindromic Subsequences
*/
func countPalindromicSubsequence(s string) int {
	first := [26]int{}
	last := [26]int{}
	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}
	for i, v := range s {
		l := v - 'a'
		last[l] = i
		if first[l] == -1 {
			first[l] = i
		}
	}
	ans := 0
	for i := 0; i < 26; i++ {
		if first[i] == -1 {
			continue
		}

		between := make(map[uint8]struct{}, 26)
		for j := first[i] + 1; j < last[i]; j++ {
			between[s[j]] = struct{}{}
		}

		ans += len(between)
	}
	return ans
}

func Test_countPalindromicSubsequence(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "",
			s:    "aabca",
			want: 3,
		},
		{
			name: "",
			s:    "adc",
			want: 0,
		},
		{
			name: "",
			s:    "bbcbaba",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPalindromicSubsequence(tt.s)
			if got != tt.want {
				t.Errorf("countPalindromicSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

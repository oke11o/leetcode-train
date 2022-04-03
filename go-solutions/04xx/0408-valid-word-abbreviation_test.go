package _4xx

import "testing"

/**
https://leetcode.com/problems/valid-word-abbreviation/
408. Valid Word Abbreviation
Easy
*/
func validWordAbbreviation(word string, abbr string) bool {
	long := 0
	short := 0
	for long < len(word) && short < len(abbr) {
		if abbr[short] == word[long] {
			short++
			long++
			continue
		}
		if abbr[short] == '0' {
			return false
		}

		shift := 0
		for abbr[short] >= '0' && abbr[short] <= '9' {
			v := abbr[short] - '0'
			if shift != 0 {
				shift *= 10
			}
			shift += int(v)
			short++
			if short == len(abbr) {
				break
			}
		}
		long += shift
		if long == len(word) || short == len(abbr) {
			break
		}
		if long > len(word) {
			return false
		}
		if abbr[short] != word[long] {
			return false
		}
	}

	return long == len(word) && short == len(abbr)
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_validWordAbbreviation(t *testing.T) {
	type args struct {
		word string
		abbr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				word: "internationalization",
				abbr: "i12iz4n",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				word: "internationalization",
				abbr: "i12iz5",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				word: "hi",
				abbr: "1",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				word: "apple",
				abbr: "a2e",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				word: "a",
				abbr: "01",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				word: "a",
				abbr: "2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validWordAbbreviation(tt.args.word, tt.args.abbr); got != tt.want {
				t.Errorf("validWordAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}

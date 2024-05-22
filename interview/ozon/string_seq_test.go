package ozon

import (
	"fmt"
	"strconv"
	"testing"
)

func StrSeq(input string) string {
	Input := []rune(input)

	var fu func(start int) (result string, finish int)

	fu = func(ptr int) (result string, finish int) {
		var mult []rune
		var substr []rune
		for {
			if isDigit(Input[ptr]) {
				if len(substr) == 0 {
					mult = append(mult, Input[ptr])
					ptr++
				} else {
					var tmp string
					tmp, ptr = fu(ptr)
					substr = append(substr, []rune(tmp)...)
				}
				continue
			}
			if isOpen(Input[ptr]) {
				ptr++
				continue
			}
			if isClose(Input[ptr]) {
				ptr++
				res := multiply(substr, mult)
				return res, ptr
			}
			if isLetter(Input[ptr]) {
				substr = append(substr, Input[ptr])
				ptr++
				continue
			}
		}
	}

	ans := ""
	i := 0
	tmp := ""
	for i < len(input) {
		tmp, i = fu(i)
		ans += tmp
	}

	return ans
}

func multiply(substr []rune, mult []rune) string {
	m, err := strconv.Atoi(string(mult))
	if err != nil {
		panic(fmt.Sprintf("invalid args(%s, %s)", string(substr), string(mult)))
	}
	var ans []rune
	for i := 0; i < m; i++ {
		for _, v := range substr {
			ans = append(ans, v)
		}
	}
	return string(ans)
}

func isDigit(v rune) bool {
	if v >= '0' && v <= '9' {
		return true
	}
	return false
}
func isOpen(v rune) bool {
	return v == '['
}
func isClose(symbol rune) bool {
	return symbol == ']'
}

func isLetter(v rune) bool {
	return !isDigit(v) && !isOpen(v) && !isClose(v)
}

func Test_StrSequence(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple",
			input: "3[a]2[bc]",
			want:  "aaabcbc",
		},
		{
			name:  "complext",
			input: "3[a]2[b3[zz]c]",
			want:  "aaabzzzzzzcbzzzzzzc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrSeq(tt.input)
			if tt.want != got {
				t.Errorf("StrSeq(`%s`) got=`%s` want=`%s`", tt.input, got, tt.want)
			}
		})
	}
}

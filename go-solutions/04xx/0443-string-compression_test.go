package _4xx

import (
	"reflect"
	"strconv"
	"testing"
)

/**
443. String Compression
Medium
https://leetcode.com/problems/string-compression/

[#two_pointers]
*/
func compress(chars []byte) int {
	prevP := 0  // поинтер, куда устанавливаем новые значения. Он же и будет конечной длинной
	firstP := 0 // поинтер, которым пробегаем по всей строке
	for firstP < len(chars) {
		curChar := chars[firstP]
		cnt := 0
		for firstP < len(chars) && chars[firstP] == curChar { // не забыть firstP < len(chars)
			firstP++
			cnt++
		}
		chars[prevP] = curChar // Нельзя пропускать, так как тут может быть все, что угодно
		prevP++
		if cnt != 1 { // тут как минимум будет 1, так как мы хоть раз зайдем в предыдущий цикл
			cntChars := strconv.Itoa(cnt)
			for i := 0; i < len(cntChars); i++ {
				chars[prevP] = cntChars[i]
				prevP++
			}
		}
	}
	return prevP
}

func Test_compress(t *testing.T) {
	tests := []struct {
		name    string
		chars   []byte
		want    int
		wantArr []byte
	}{
		{
			name:    "",
			chars:   []byte("aabbccc"),
			want:    6,
			wantArr: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			name:    "",
			chars:   []byte("a"),
			want:    1,
			wantArr: []byte{'a'},
		},
		{
			name:    "",
			chars:   []byte("abbbbbbbbbbbb"),
			want:    4,
			wantArr: []byte{'a', 'b', '1', '2'},
		},
		{
			name:    "",
			chars:   []byte("aaabbaa"),
			want:    4,
			wantArr: []byte{'a', '3', 'b', '2', 'a', '2'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compress(tt.chars)
			if got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
			tt.chars = tt.chars[0:got]
			if !reflect.DeepEqual(tt.chars, tt.wantArr) {
				t.Errorf("compress() = %v, wantArr %v", string(tt.chars), string(tt.wantArr))
			}
		})
	}
}

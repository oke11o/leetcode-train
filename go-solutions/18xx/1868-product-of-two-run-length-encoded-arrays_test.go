package _8xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/product-of-two-run-length-encoded-arrays/
1868. Product of Two Run-Length Encoded Arrays
Medium

Обратить внимание:
- тут сделано без изменения входных слайсов. С изменением проще
- нужно вычитать и из предыдущей частоты. Очень запутано
*/
func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	result := [][]int{}
	if len(encoded1) == 0 {

		return result
	}
	i1 := 0
	i2 := 0

	freq1 := encoded1[i1][1]
	freq2 := encoded2[i2][1]
	for i1 < len(encoded1) || i2 < len(encoded2) {
		val1 := encoded1[i1][0]
		val2 := encoded2[i2][0]
		if !(i1 == 0 && i2 == 0) {
			if freq1 == freq2 {
				freq1 = encoded1[i1][1]
				freq2 = encoded2[i2][1]
			} else if freq1 < freq2 {
				freq2 -= freq1
				freq1 = encoded1[i1][1]
			} else {
				freq1 -= freq2
				freq2 = encoded2[i2][1]
			}
		}

		res := val1 * val2
		freq := freq1
		if freq > freq2 {
			freq = freq2
		}
		if len(result) != 0 && res == result[len(result)-1][0] {
			result[len(result)-1][1] += freq
		} else {
			result = append(result, []int{res, freq})
		}

		if freq1 == freq2 {
			i1++
			i2++
		} else if freq1 < freq2 {
			i1++
		} else {
			i2++
		}
	}
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_findRLEArray(t *testing.T) {
	type args struct {
		encoded1 [][]int
		encoded2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				encoded1: [][]int{{1, 3}, {2, 3}},
				encoded2: [][]int{{6, 3}, {3, 3}},
			},
			want: [][]int{{6, 6}},
		},
		{
			name: "",
			args: args{
				encoded1: [][]int{{1, 3}, {2, 1}, {3, 2}},
				encoded2: [][]int{{2, 3}, {3, 3}},
			},
			want: [][]int{{2, 3}, {6, 1}, {9, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRLEArray(tt.args.encoded1, tt.args.encoded2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRLEArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

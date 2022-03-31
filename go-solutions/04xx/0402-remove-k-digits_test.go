package _4xx

import "testing"

/**
https://leetcode.com/problems/remove-k-digits/
402. Remove K Digits
Medium
#stack, #greedy, #string, #facebook

Суть в том, что мы должны делать возрастающий (не убывающий) моностек.
То есть выкидывать все предыдущие значения, если нашли меньшее число. Но делать это всего k раз.
*/
func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	monoStack := make([]byte, 0, k)
	for i := 0; i < len(num); i++ {
		// Суть в том, что мы должны делать возрастающий (не убывающий) моностек.
		// То есть выкидывать все предыдущие значения, если нашли меньшее число. Но делать это всего k раз.
		for len(monoStack) > 0 && k > 0 && monoStack[len(monoStack)-1] > num[i] {
			monoStack = monoStack[:len(monoStack)-1]
			k--
		}

		monoStack = append(monoStack, num[i])
	}

	// remove the remaining digits from the tail.
	monoStack = monoStack[:len(monoStack)-k]

	// build the final string, while removing the leading zeros.
	leadingZero := true
	result := make([]byte, 0)
	for i := 0; i < len(monoStack); i++ {
		if leadingZero && monoStack[i] == '0' {
			continue
		}
		leadingZero = false
		result = append(result, monoStack[i])
	}

	if len(result) == 0 {
		return "0"
	}

	return string(result)
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				num: "1432219",
				k:   3,
			},
			want: "1219",
		},
		{
			name: "",
			args: args{
				num: "10",
				k:   2,
			},
			want: "0",
		},
		{
			name: "",
			args: args{
				num: "10",
				k:   1,
			},
			want: "0",
		},
		{
			name: "",
			args: args{
				num: "10200",
				k:   1,
			},
			want: "200",
		},
		{
			name: "",
			args: args{
				num: "112",
				k:   1,
			},
			want: "11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

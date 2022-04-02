package _4xx

import "testing"

/**
https://leetcode.com/problems/remove-k-digits/
402. Remove K Digits
Medium
#stack, #greedy, #string, #facebook

Brute force.
1. Helper func for compare string with same length. Just comparing letters from string start. And first not equal letter is our result.
12 3 412  smaller
12 5 412  bigger
2. Brute force approch. Use stratagy for removing k elements. And on each iteration of removing we should compare string. And store the smallest.
3. What strategy? Using permutation algo. 123 412, 124 312, 121 342...
4. But time complexity is really, really high. It is factorial of length of string

The another approch - is use monotonic stack. We can represent input string as subsequnce of monotonic increasing and monotonic decreasing sequeces.
1234312
31
43
23412 - bigger.
13412

1432219
4>3 - should remove


Суть в том, что мы должны делать возрастающий (не убывающий) моностек.
То есть выкидывать все предыдущие значения, если нашли меньшее число. Но делать это всего k раз.
*/
func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	result := []byte{}
	// convert input to custom monotonic stack
	for i := 0; i < len(num); i++ {
		// Суть в том, что мы должны делать возрастающий (не убывающий) моностек.
		// То есть выкидывать все предыдущие значения, если нашли меньшее число. Но делать это всего k раз.
		for k > 0 && len(result) > 0 && result[len(result)-1] > num[i] {
			result = result[0 : len(result)-1]
			k--
		}
		result = append(result, num[i])
	}
	// if input monotonic increasing more than k times
	// remove the remaining digits from the tail.
	result = result[0 : len(result)-k]

	// remove leading zeroes
	for len(result) > 0 && result[0] == '0' {
		result = result[1:]
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

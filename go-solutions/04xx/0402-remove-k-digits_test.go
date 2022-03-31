package _4xx

import "testing"

/**
https://leetcode.com/problems/remove-k-digits/
402. Remove K Digits
Medium
#stack, #greedy, #string, #facebook
*/
func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	monoStack := make([]byte, 0, k)
	for i := 0; i < len(num); i++ {
		for len(monoStack) > 0 && monoStack[len(monoStack)-1] < num[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}

		monoStack = append(monoStack, num[i])
		if len(monoStack) == k {
			break
		}
	}
	k = k - len(monoStack) //обнулю K. Зачему? Чтобы потом обрезать с конца строку, если убывающий символов не достаточно
	// сконвертирую stack 2 map. Чтобы было проще проверять и удалять.
	m := make(map[byte]int, len(monoStack))
	for _, v := range monoStack {
		m[v]++
	}
	// Добавлю в результат только те, кого нет в мапе
	result := make([]byte, 0, len(num)-len(m))
	for i := 0; i < len(num); i++ {
		digit := num[i]
		if _, ok := m[digit]; ok {
			m[digit]--
			if m[digit] == 0 {
				delete(m, digit)
			}
		} else {
			if !(len(result) == 0 && digit == '0') {
				result = append(result, digit)
			}
		}
	}

	// Если убывающий символов недостаточно, обрезу
	result = result[0 : len(result)-k]
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

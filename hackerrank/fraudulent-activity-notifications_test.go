package hackerrank

import (
	"sort"
	"testing"
)

/**
1. Первая мысль. (Неверная) Отсортировать, и потом просто считать median, как смещение назад на половину подмассива.
Это не верный подход. Так как отсортированный моссив может быть плавно возрастающим, но большим.
А большие числа могут появляться в середине.
То есть при плавно возрастающем отсортированном массиве, мы получим, что никогда не надо делать нотификации.
Но на самом деле могут быть пики, когда надо делать.
------
2. Попробовать брутфорсом. ИИИИИ. Не проходим по времени! Ищу решение в инете. И там интересный подход. С массивом в 201 элемент.
Что?! Смотрю наши constraints. И вижу, что 0<=expenditure[i]<=200
Пробую этот вариант
*/

/**
https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem
Fraudulent Activity Notifications
Medium
*/
func activityNotifications(expenditure []int32, D int32) int32 {
	var findMid = func(values [201]int, days int) float64 {
		sum := 0
		var mid1, mid2 float64
		for i := 0; i < 201; i++ {
			sum += values[i]
			if sum <= days/2 {
				mid1 = 0

			}
		}
		_, _ = mid1, mid2
		return 0
	}
	_ = findMid
	d := int(D)

	if len(expenditure) < d+1 {
		return 0
	}
	offset := d / 2
	isEven := d%2 == 0

	values := [201]int{}
	for i := 0; i < d; i++ {
		values[expenditure[i]]++
	}

	var result int32
	for i := d; i < len(expenditure); i++ {
		subArr := make([]int32, d)
		copy(subArr, expenditure[i-d:i])
		sort.Slice(subArr, func(i, j int) bool {
			return subArr[i] < subArr[j]
		})
		median := float64(subArr[d-offset-1])
		if isEven {
			median = (median + float64(expenditure[d-offset])) / 2
		}
		// get Result
		if float64(expenditure[i]) >= 2*median {
			result++
		}
	}

	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_activityNotifications(t *testing.T) {
	type args struct {
		expenditure []int32
		d           int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "",
			args: args{
				expenditure: []int32{2, 3, 4, 2, 3, 6, 8, 4, 5},
				d:           5,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				expenditure: []int32{1, 2, 3, 4, 4},
				d:           4,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				expenditure: []int32{1, 2, 3, 4, 5},
				d:           3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := activityNotifications(tt.args.expenditure, tt.args.d); got != tt.want {
				t.Errorf("activityNotifications() = %v, want %v", got, tt.want)
			}
		})
	}
}

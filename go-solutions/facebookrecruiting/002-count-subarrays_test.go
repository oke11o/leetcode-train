package facebookrecruiting

import (
	"reflect"
	"testing"
)

/**
Contiguous Subarrays
You are given an array arr of N integers. For each index i, you are required to determine the number of contiguous subarrays that fulfill the following conditions:
The value at index i must be the maximum element in the contiguous subarrays, and
These contiguous subarrays must either start from or end on index i.
Signature
int[] countSubarrays(int[] arr)
Input
Array arr is a non-empty list of unique integers that range between 1 to 1,000,000,000
Size N is between 1 and 1,000,000
Output
An array where each index i contains an integer denoting the maximum number of contiguous subarrays of arr[i]
Example:
arr = [3, 4, 1, 6, 2]
output = [1, 3, 1, 5, 1]
Explanation:
For index 0 - [3] is the only contiguous subarray that starts (or ends) with 3, and the maximum value in this subarray is 3.
For index 1 - [4], [3, 4], [4, 1]
For index 2 - [1]
For index 3 - [6], [6, 2], [1, 6], [4, 1, 6], [3, 4, 1, 6]
For index 4 - [2]
So, the answer for the above input is [1, 3, 1, 5, 1]
*/
// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
func countSubarrays_bruteForce(arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		for j := i; j >= 0 && arr[j] <= v; j-- {
			result[i]++
		}
		for j := i; j < len(arr) && arr[j] <= v; j++ {
			result[i]++
		}
		result[i]--
	}
	return result
}
func countSubarrays(arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		for j := i; j >= 0 && arr[j] <= v; j-- {
			result[i]++
		}
		for j := i; j < len(arr) && arr[j] <= v; j++ {
			result[i]++
		}
		result[i]--
	}
	return result
}

// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
//15,14,16,11,12,13,12,11,10, 9,
//-1, 1,-1, 2, 2, 2, 5, 6, 7, 8
//               1  0
//Изначально у всех 1
//1, 1, 1, 1, 1, 1
//2, 1, 1, 1, 1, 1
//3, 2, 1, 1, 1, 1
//1 -- сколько элементов слева больше чем arr[i]?
//  ^ -- 2 больше 3? - да. У 3 добавляем 1. У 2 ставим 0
//     ^ -- 4 больше 3? - нет. Закончили с 3.

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_countSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				arr: []int{3, 4, 1, 6, 2},
			},
			want: []int{1, 3, 1, 5, 1},
		},
		{
			name: "",
			args: args{
				arr: []int{2, 4, 7, 1, 5, 3},
			},
			want: []int{1, 2, 6, 1, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

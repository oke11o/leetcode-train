package facebookrecruiting

import (
	"testing"
)

/**
https://www.facebookrecruiting.com/portal/coding_practice_question/?problem_id=238827593802550&ppid=454615229006519&practice_plan=1
Rotational Cipher
One simple way to encrypt a string is to "rotate" every alphanumeric character by a certain amount. Rotating a character means replacing it with another character that is a certain number of steps away in normal alphabetic or numerical order.
For example, if the string "Zebra-493?" is rotated 3 places, the resulting string is "Cheud-726?". Every alphabetic character is replaced with the character 3 letters higher (wrapping around from Z to A), and every numeric character replaced with the character 3 digits higher (wrapping around from 9 to 0). Note that the non-alphanumeric characters remain unchanged.
Given a string and a rotation factor, return an encrypted string.
Signature
string rotationalCipher(string input, int rotationFactor)
Input
1 <= |input| <= 1,000,000
0 <= rotationFactor <= 1,000,000
Output
Return the result of rotating input a number of times equal to rotationFactor.
Example 1
input = Zebra-493?
rotationFactor = 3
output = Cheud-726?
Example 2
input = abcdefghijklmNOPQRSTUVWXYZ0123456789
rotationFactor = 39
output = nopqrstuvwxyzABCDEFGHIJKLM9012345678
*/
// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
func rotationalCipher(input string, rotationFactor int) string {
	if len(input) == 0 || rotationFactor == 0 {
		return input
	}
	result := make([]rune, len(input))
	for i, v := range input {
		var left, right rune
		if v >= '0' && v <= '9' {
			left, right = '0', '9'
		} else if v >= 'A' && v <= 'Z' {
			left, right = 'A', 'Z'
		} else if v >= 'a' && v <= 'z' {
			left, right = 'a', 'z'
		}
		if right != 0 {
			v -= left
			v += rune(rotationFactor)
			v = v % (right - left + 1)
			v += left
		}
		result[i] = v
	}
	// Write your code here
	return string(result)
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_rotationalCipher(t *testing.T) {
	type args struct {
		input          string
		rotationFactor int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				input:          "Zebra-493?",
				rotationFactor: 3,
			},
			want: "Cheud-726?",
		},
		{
			name: "",
			args: args{
				input:          "abcdefghijklmNOPQRSTUVWXYZ0123456789",
				rotationFactor: 39,
			},
			want: "nopqrstuvwxyzABCDEFGHIJKLM9012345678",
		},
		{
			name: "",
			args: args{
				input:          "abcDEF123",
				rotationFactor: 1,
			},
			want: "bcdEFG234",
		},
		{
			name: "",
			args: args{
				input:          "0aA123abcDEF",
				rotationFactor: 9,
			},
			want: "9jJ012jklMNO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotationalCipher(tt.args.input, tt.args.rotationFactor); got != tt.want {
				t.Errorf("rotationalCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}

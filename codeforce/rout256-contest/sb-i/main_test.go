package main

import "testing"

func Test_problemI(t *testing.T) {
	tests := []struct {
		name string
		word string
		dict []string
		want string
	}{
		{
			name: "",
			word: "flask",
			dict: []string{"task", "decide", "id"},
			want: "task",
		},
		{
			name: "",
			word: "code",
			dict: []string{"task", "decide", "id"},
			want: "decide",
		},
		{
			name: "",
			word: "void",
			dict: []string{"task", "decide", "id"},
			want: "id",
		},
		{
			name: "",
			word: "forces",
			dict: []string{"task", "decide", "id"},
			want: "task",
		},
		{
			name: "",
			word: "id",
			dict: []string{"task", "decide", "id"},
			want: "task",
		},
		{
			name: "",
			word: "aauaua",
			dict: []string{"ua", "u", "uauaauuu", "auauauua", "aauuau"},
			want: "ua",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewTrie()
			for i, d := range tt.dict {
				trie.Insert(reverseString(d), i)
			}
			if got := problemI(tt.word, tt.dict, trie); got != tt.want {
				t.Errorf("problemI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findZarif(t *testing.T) {
	tests := []struct {
		name string
		word string
		d    string
		want int
	}{
		{
			name: "",
			word: "task",
			d:    "flask",
			want: 3,
		},
		{
			name: "",
			word: "decide",
			d:    "code",
			want: 2,
		},
		{
			name: "",
			word: "id",
			d:    "void",
			want: 2,
		},
		{
			name: "",
			word: "code",
			d:    "forces",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findZarif(tt.word, tt.d); got != tt.want {
				t.Errorf("findZarif() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "",
			in:   "qwer",
			want: "rewq",
		},
		{
			name: "",
			in:   "asdfg",
			want: "gfdsa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.in); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

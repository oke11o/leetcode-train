package main

import "testing"

func Test_problemI(t *testing.T) {
	tests := []struct {
		name string
		word string
		dict []string
		want string
	}{
		{name: "", word: "vwlldlswl", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "svwl"},
		{name: "", word: "slvlvdldw", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "svdw"},
		{name: "", word: "vwls", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "lvvlvsdss"},
		{name: "", word: "dl", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "vwswll"},
		{name: "", word: "lvs", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "sss"},
		{name: "", word: "ldvsl", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "svwl"},
		{name: "", word: "ddlsw", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "wsvvlsw"},
		{name: "", word: "d", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "ddvlwswlsd"},
		{name: "", word: "svsvvwsdwl", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "svwl"},
		{name: "", word: "vlssldswlw", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "wsvvlsw"},
		{name: "", word: "lvlsdl", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "dl"},
		{name: "", word: "s", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "wslswds"},
		{name: "", word: "vlwlssld", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "ddvlwswlsd"},
		{name: "", word: "ldvl", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "svwl"},
		{name: "", word: "dvvvwsssww", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "swvw"},
		{name: "", word: "vdlvldw", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "svdw"},
		{name: "", word: "wwv", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "ldlddvddvv"},
		{name: "", word: "dlwlwwsvdl", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "dl"},
		{name: "", word: "swsw", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "wsvvlsw"},
		{name: "", word: "dwldddwwl", dict: []string{"sss", "ddvlwswlsd", "swvw", "dl", "llwllvw", "swdsdwsvd", "vwvsv", "wslswds", "vwswll", "wwvdsslwd", "wsvvlsw", "lvvlvsdss", "ldwlvsd", "v", "wvldwd", "ldlddvddvv", "vsswwll", "svwl", "sldwdsvsv", "svdw"}, want: "svwl"},

		{name: "", word: "v", dict: []string{"ve", "neeevzzv"}, want: "neeevzzv"},
		{name: "", word: "vnveneezve", dict: []string{"ve", "neeevzzv"}, want: "ve"},
		{name: "", word: "evzz", dict: []string{"ve", "neeevzzv"}, want: "neeevzzv"},
		{name: "", word: "ezenne", dict: []string{"ve", "neeevzzv"}, want: "ve"},
		{name: "", word: "nneenn", dict: []string{"ve", "neeevzzv"}, want: "neeevzzv"},
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
			word: "uuauuauuau",
			dict: []string{"ua", "u", "uauaauuu", "auauauua", "aauuau"},
			want: "aauuau",
		},
		{
			name: "",
			word: "aaau",
			dict: []string{"ua", "u", "uauaauuu", "auauauua", "aauuau"},
			want: "aauuau",
		},
		{
			name: "",
			word: "a",
			dict: []string{"ua", "u", "uauaauuu", "auauauua", "aauuau"},
			want: "ua",
		},
		{
			name: "",
			word: "uuuaaauaaa",
			dict: []string{"ua", "u", "uauaauuu", "auauauua", "aauuau"},
			want: "ua",
		},
		{
			name: "",
			word: "aauaua",
			dict: []string{"ua", "u", "uauaauuu", "auauauua", "aauuau"},
			want: "ua", // "auauauua",
		},
		{
			name: "",
			word: "uaaa",
			dict: []string{"ua", "u", "uauaauuu", "auauauua", "aauuau"},
			want: "ua",
		},
		{
			name: "",
			word: "auuuu",
			dict: []string{"ua", "u", "uauaauuu", "auauauua", "aauuau"},
			want: "uauaauuu",
		},
		{
			name: "",
			word: "auuaaauu",
			dict: []string{"ua", "u", "uauaauuu", "auauauua", "aauuau"},
			want: "uauaauuu",
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

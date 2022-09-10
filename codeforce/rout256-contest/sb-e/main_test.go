package main

import "testing"

func Test_parseAndValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    int
		wantErr bool
	}{
		{
			name:    "",
			s:       "02:46:00",
			want:    9960,
			wantErr: false,
		},
		{
			name:    "",
			s:       "24:00:00",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseAndValidate(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseAndValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseAndValidate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_problemE(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{
			name: "",
			arr:  []string{"02:46:00-03:14:59"},
			want: "YES",
		},
		{
			name: "",
			arr:  []string{"23:59:59-23:59:59", "00:00:00-23:59:58"},
			want: "YES",
		},
		{
			name: "",
			arr:  []string{"23:59:58-23:59:59", "00:00:00-23:59:58"},
			want: "NO",
		},
		{
			name: "",
			arr:  []string{"23:59:59-23:59:58", "00:00:00-23:59:57"},
			want: "NO",
		},
		{
			name: "",
			arr:  []string{"17:53:39-20:20:02", "10:39:17-11:00:52", "08:42:47-09:02:14", "09:44:26-10:21:41", "00:46:17-02:07:19", "22:42:50-23:17:46"},
			want: "YES",
		},
		{
			name: "",
			arr:  []string{"24:00:00-23:59:59"},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problemE(tt.arr); got != tt.want {
				t.Errorf("problemE() = %v, want %v", got, tt.want)
			}
		})
	}
}

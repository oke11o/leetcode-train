package main

import "testing"

func Test_parseDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		want1   int64
		wantErr bool
	}{
		{
			name: "",
			args: args{
				date: "00:02:58,225 --> 00:03:00,561",
			},
			want:    178225,
			want1:   180561,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				date: "10:00:00,000 --> 10:00:00,009",
			},
			want:    36000000,
			want1:   36000009,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseDate(tt.args.date)
			if (tt.wantErr && err == nil) || (!tt.wantErr && err != nil) {
				t.Errorf("%+v", err)
			}
			if got != tt.want {
				t.Errorf("parseDate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseDate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_dates2date(t *testing.T) {
	type args struct {
		from int64
		to   int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				from: 178225,
				to:   180561,
			},
			want: "00:02:58,225 --> 00:03:00,561",
		},
		{
			name: "",
			args: args{
				from: 36000000,
				to:   36000009,
			},
			want: "10:00:00,000 --> 10:00:00,009",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dates2date(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("dates2date() = %v, want %v", got, tt.want)
			}
		})
	}
}

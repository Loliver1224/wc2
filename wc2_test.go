package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_count(t *testing.T) {
	type args struct {
		file *os.File
	}
	test1, _ := os.Open("testfiles/alice_win.txt")
	test2, _ := os.Open("testfiles/alice_mac.txt")
	test3, _ := os.Open("testfiles/alice_unix.txt")
	test4, _ := os.Open("testfiles/multibyte_chars.txt")
	tests := []struct {
		name    string
		args    args
		wantCnt *Counter
	}{
		{
			name:    "CR+LF test",
			args:    args{test1},
			wantCnt: &Counter{603, 583, 11, 72, 112},
		},
		{
			name:    "CR test",
			args:    args{test2},
			wantCnt: &Counter{593, 583, 11, 72, 112},
		},
		{
			name:    "LF test",
			args:    args{test3},
			wantCnt: &Counter{593, 583, 11, 72, 112},
		},
		{
			name:    "multibyte-chars test",
			args:    args{test4},
			wantCnt: &Counter{112, 35, 3, 13, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCnt := count(tt.args.file); !reflect.DeepEqual(gotCnt, tt.wantCnt) {
				t.Errorf("count() = %v, want %v", gotCnt, tt.wantCnt)
			}
		})
	}
}

func Test_isBreak(t *testing.T) {
	type args struct {
		prev rune
		curr rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nbsp",
			args: args{'a', 'b'},
			want: false,
		},
		{
			name: "alpha + CR",
			args: args{'y', '\r'},
			want: true,
		},
		{
			name: "alpha + LF",
			args: args{'y', '\n'},
			want: true,
		},
		{
			name: "CR + LF",
			args: args{'\r', '\n'},
			want: true,
		},
		{
			name: "LF * 2",
			args: args{'\n', '\n'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBreak(tt.args.prev, tt.args.curr); got != tt.want {
				t.Errorf("isBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "file exist",
			args: args{"wc2.go"},
			want: true,
		},
		{
			name: "file not exist",
			args: args{"gopher.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFileExists(tt.args.filename); got != tt.want {
				t.Errorf("isFileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

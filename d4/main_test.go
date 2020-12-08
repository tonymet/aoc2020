package main

import (
	"testing"
)

func Test_checkPid(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"passport",
			args{"123456789"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPid(tt.args.text); got != tt.want {
				t.Errorf("checkPid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkEye(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"eye",
			args{"brn"},
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEye(tt.args.text); got != tt.want {
				t.Errorf("checkEye() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkYear(t *testing.T) {
	type args struct {
		test string
		min  int64
		max  int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1920-2002",
			args{"1930", 1920, 2002},
			true,
		}, // TODO: Add test cases.
		{
			"1920-2002",
			args{"1910", 1920, 2002},
			false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkYear(tt.args.test, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("checkYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkHeight(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"height",
			args{"150cm"},
			true,
		},
		{
			"height",
			args{"400cm"},
			false,
		},
		{
			"height",
			args{"77in"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkHeight(tt.args.text); got != tt.want {
				t.Errorf("checkHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

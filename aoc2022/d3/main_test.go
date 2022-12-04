package main

import (
	"testing"
)

func Test_charVal(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test a",
			args{'a'},
			1},
		{"test A",
			args{'A'},
			27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := charVal(tt.args.c); got != tt.want {
				t.Errorf("charVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

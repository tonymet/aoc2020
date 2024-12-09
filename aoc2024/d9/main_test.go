package main

import (
	"io"
	"reflect"
	_ "sort"
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{strings.NewReader("12345")},
			want: "0..111....22222",
		},
		{
			name: "test2",
			args: args{strings.NewReader("2333133121414131402")},
			want: "00...111...2...333.44.5555.6666.777.888899",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInput(tt.args.in); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

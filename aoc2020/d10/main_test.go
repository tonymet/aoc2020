package main

import (
	"reflect"
	"testing"
)

func Test_findChains(t *testing.T) {
	type args struct {
		records recordType
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			"example",
			args{recordType{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4, 0, 22}},
			map[int]int{1: 7, 3: 5},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findChains(tt.args.records)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkTargetSum() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			findChains(tt.args.records)
		})
	}
}

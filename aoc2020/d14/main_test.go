package main

import "testing"

func Test_memValue(t *testing.T) {
	type args struct {
		inboundValue  int64
		outboundValue int64
		trueMask      int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"11",
			args{11, 64, 68719476669},
			73,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := memValue(tt.args.inboundValue, tt.args.outboundValue, tt.args.trueMask); got != tt.want {
				t.Errorf("memValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_divCmpInt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    int64
		want bool
	}{
		{"test 222221	", 222221, false},
		{"test 97979797	", 97979797, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divCmpInt(tt.a)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("divCmpInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_divCmpInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divCmpInt(97979797)
	}
}

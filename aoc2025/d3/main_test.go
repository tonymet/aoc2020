package main

import (
	"testing"
)

func Test_findMax3(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
		v    []int64
		s    string
		want int64
	}{
		{"test1", []int64{}, "6739459674389333459433695375559949344734767926833587823236783998689734978783695374574455875833736627", 5544},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMax3(tt.v)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("findMax3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findMax3(b *testing.B) {
	v, _ := stringInts("6739459674389333459433695375559949344734767926833587823236783998689734978783695374574455875833736627")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMax3(v)
	}
}

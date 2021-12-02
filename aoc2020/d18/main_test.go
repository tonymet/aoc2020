package main

import "testing"

func Test_istack_push(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		is   *istack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.is.push(tt.args.a)
		})
	}
}

func BenchmarkPushPop(b *testing.B) {
	// run the Fib function b.N times
	var a istack
	for n := 0; n < b.N; n++ {
		a.push(5)
		a.pop()
	}
}

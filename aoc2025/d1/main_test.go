package main

import (
	"reflect"
	"testing"
)

func Test_dial_op(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r      Rot
		d      dial
		e      dial
		wantZc int64
	}{
		{"test1", Rot{dir: 'L', len: 898}, dial{pos: 89}, dial{91}, 9},
		//d.pos : 0       , rot : L224    d.pos : 76, d.zc : 2
		{"test2", Rot{dir: 'L', len: 224}, dial{pos: 0}, dial{76}, 2},
		// d.pos : 99      , rot : L530    d.pos : 69, d.zc : 4
		{"test L530", Rot{dir: 'L', len: 530}, dial{pos: 99}, dial{69}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zc := tt.d.op(tt.r)
			if !(reflect.DeepEqual(tt.d, tt.e) && zc == tt.wantZc) {
				t.Logf("got: %x, %d, want %x, %d", tt.d, zc, tt.e, tt.wantZc)
				t.FailNow()
			}
		})
	}
}

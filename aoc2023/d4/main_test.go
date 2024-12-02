package main

import (
	_ "sort"
	"strings"
	"testing"
)

func Test_rec_solveY(t *testing.T) {
	r1 := parseRec(strings.NewReader("19, 13, 30 @ -2,  1, -2"))
	r2 := parseRec(strings.NewReader("18, 19, 22 @ -1, -1, -2"))
	y1 := r1.solveY(14.333)
	y2 := r2.solveY(14.333)
	if y1 != y2 {
		t.Logf("y1 = %f, y2=%f", y1, y2)
		t.Fail()
	}
}

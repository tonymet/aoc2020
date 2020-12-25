package main

import (
	"reflect"
	"testing"
)

func Test_DeepEqual(t *testing.T) {
	if !reflect.DeepEqual([]int{1, 2, 5, 10}, []int{1, 2, 5, 10}) {
		t.Errorf("deepEqual doesn't match\n")
	}
}

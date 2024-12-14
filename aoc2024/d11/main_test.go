package main

import (
	"os"
	_ "sort"
	"testing"
)

func Benchmark_Stoners(b *testing.B) {
	f, err := os.Open("data/input.txt")
	part = 2
	if err != nil {
		panic(err)
	}
	for range b.N {
		f.Seek(0, 0)
		stoners(f)
	}
}

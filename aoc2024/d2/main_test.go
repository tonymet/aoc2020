package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func BenchmarkPart2(b *testing.B) {
	// run the Fib function b.N times
	silent = true
	part = 2
	var buf bytes.Buffer
	f, err := os.Open("data/aoc2024d2.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(&buf, f)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		reader := bytes.NewReader(buf.Bytes())
		parseFile(reader)
	}
}

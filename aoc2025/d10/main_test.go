package main

import "testing"

func Test_testSolution(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		target  []rune
		buttons [][]int
		want    bool
	}{
		{".##.", []rune{'.', '#', '#', '.'}, [][]int{{0, 2}, {0, 1}}, true},
		{"...#.", []rune{'.', '.', '.', '#', '.'}, [][]int{{0, 4}, {0, 1, 2}, {1, 2, 3, 4}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testSolution(tt.target, tt.buttons)
			if got != tt.want {
				t.Errorf("testSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

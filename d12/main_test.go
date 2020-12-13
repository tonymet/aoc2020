package main

import "testing"

func Test_newBearing(t *testing.T) {
	type args struct {
		curBearing string
		turnCmd    cmd
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"R90",
			args{"N", cmd{"R", 90}},
			"E",
		},
		{
			"L90",
			args{"N", cmd{"L", 90}},
			"W",
		},
		{
			"L90",
			args{"E", cmd{"L", 90}},
			"N",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBearing(tt.args.curBearing, tt.args.turnCmd); got != tt.want {
				t.Errorf("newBearing() = %v, want %v", got, tt.want)
			}
		})
	}
}

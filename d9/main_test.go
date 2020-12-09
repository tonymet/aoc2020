package main

import "testing"

/*

	20
	15
	25
	47
	40
	62
	55
	65
	95
	102
	117
	150
	182
	127
	219
	299
	277
	309
	576}

*/

func Test_checkTargetSum(t *testing.T) {
	type args struct {
		target int
		pool   recordType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"40",
			args{40, recordType{35, 20, 15, 25, 47}},
			true,
		},
		{
			"127",
			args{127, recordType{95, 102, 117, 150, 182}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkTargetSum(tt.args.target, tt.args.pool); got != tt.want {
				t.Errorf("checkTargetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

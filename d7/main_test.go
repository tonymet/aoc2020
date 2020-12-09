package main

import (
	"testing"
)

func Test_scanLine(t *testing.T) {
}

func Test_countChildBags(t *testing.T) {
	type args struct {
		cur Content
	}
	tests := []struct {
		name string
		args Content
		want uint64
	}{
		{"test1", Content{"shiny gold", 1}, 15},
	}
	bagIndex.parentChildContents = make(map[string][]Content)
	bagIndex.parentChildContents["shiny gold"] = []Content{{"red", 2}}
	bagIndex.parentChildContents["red"] = []Content{{"green", 4}, {"yellow", 2}}
	//bagIndex.parentChildContents["yellow"] = []Content{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countChildBags(tt.args); got != tt.want {
				t.Errorf("countChildBags() = %v, want %v", got, tt.want)
			}
		})
	}
}

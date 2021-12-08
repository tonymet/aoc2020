use Test::More tests => 1;
is(`cat data/input.txt| PART=1 go run main.go|grep "board id" | tail -n 1`, "board id: 100\n")

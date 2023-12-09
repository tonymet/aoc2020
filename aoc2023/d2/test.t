use Test::More tests => 2; # or
is((system "cat data/testinput.txt| go run . | grep 'total p2 : 2286'"), 0, "2286 expected");
is(`cat data/testinput.txt| go run . | grep 'total p2'`, "total p2 : 2286\n", "2286 expected");